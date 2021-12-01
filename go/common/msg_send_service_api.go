package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type msgSendFunc func(string, string) error

//神州浩天企业微信消息发送Info
var corpid = "wxa059996e5d72516b"
var corpsecret = "DeLr5iZm4-bicImGQG344kfUDt2jer8tg-iIYm9cxwA"
var agent_id = "1000008"
var url_getToken = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=` + corpid + `&corpsecret=` + corpsecret
var url_postSend = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=`

//神州浩天企业微信消息发送
var MsgSendFunc_Szht msgSendFunc = func(toUserId string, msg string) error {

	resp, err := http.Get(url_getToken)
	defer resp.Body.Close()
	if err != nil {
		return errors.New("无法请求微信获取tokken::" + err.Error())
	}
	contentByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("读取微信返回值失败::" + err.Error())
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(contentByte, &jsonMap)
	if err != nil {
		return errors.New("微信返回数据解析错误::" + string(contentByte) + "||" + err.Error())
	}

	var access_token = jsonMap["access_token"].(string)

	userId := toUserId

	postSendMsgMap := make(map[string]interface{})
	postSendMsgMap["touser"] = userId
	postSendMsgMap["msgtype"] = "text"
	postSendMsgMap["agentid"] = agent_id
	postSendMsgMap["safe"] = "0"
	contentMap := make(map[string]interface{})

	contentMap["content"] = msg
	postSendMsgMap["text"] = contentMap
	postSendMsgBytes, err := json.Marshal(postSendMsgMap)
	if err != nil {
		//这个错误不会出错~！
	}
	respPost, errPost := http.Post(url_postSend+access_token, "application/json", bytes.NewReader(postSendMsgBytes))
	defer respPost.Body.Close()
	if errPost != nil {
		return errors.New("无法请求微信发送::" + errPost.Error())
	}
	contentBytePost, errPost := ioutil.ReadAll(respPost.Body)
	if errPost != nil {
		return errors.New("读取微信返回值失败::" + errPost.Error())
	}
	jsonMapPost := make(map[string]interface{})
	err = json.Unmarshal(contentBytePost, &jsonMapPost)
	if err != nil {
		return errors.New("微信返回数据解析错误::" + string(contentBytePost) + "||" + err.Error())
	}
	return nil
}

//清华大学企业微信号推送
var qhdx_pushmsg_key = "cwxxfw"     //微信服务接口账号
var qhdx_pushmsg_value = "LWZch0SV" //微信服务接口密码

type qhdxPostStruct struct {
	Errorcode   int    `json:"errorcode"`
	Invaliduser string `json:"invaliduser"`
}

//清华大学企业微信消息发送
var MsgSendFunc_Qhdx msgSendFunc = func(toUserId string, msg string) error {
	//清华大学
	teacherCode := toUserId
	get_resp, err := http.Get("http://weixin.cic.tsinghua.edu.cn/cop/getTimestamp.php")
	defer get_resp.Body.Close()
	if err != nil {
		return errors.New("获取getTimestamp失败：：" + err.Error())
	}
	b, _ := ioutil.ReadAll(get_resp.Body)
	timestamp := string(b)
	key := StringToMd5(qhdx_pushmsg_key + timestamp + qhdx_pushmsg_value)
	postUrl := "http://weixin.cic.tsinghua.edu.cn/cop/sendmsg.php?app=" + qhdx_pushmsg_key + "&key=" + key +
		"&timestamp=" + timestamp + "&type=text"
	//var postData = "userlist=" + URLEncoder.encode(userList, "utf-8") + "&safe=0&msg=" + URLEncoder.encode(message, "utf-8");
	postData := ""
	postData = postData + "userlist=" + url.QueryEscape(teacherCode) + "&safe=0&msg=" + url.QueryEscape(msg)
	resp, err := http.Post(postUrl, "application/x-www-form-urlencoded", strings.NewReader(postData))
	defer resp.Body.Close()
	if err != nil {
		return errors.New("清华高校负责人消息发送失败：：" + err.Error())
	}
	re, _ := ioutil.ReadAll(resp.Body)
	qhdxPostStruct := new(qhdxPostStruct)
	errUn := json.Unmarshal(re, qhdxPostStruct)
	if errUn != nil {
		return errors.New("解码qhdxPostStruct返回值错误失败：：" + errUn.Error())
	}
	if qhdxPostStruct.Errorcode != 0 {
		return errors.New("qhdxPostStruct.Errorcode失败：：" + strconv.Itoa(qhdxPostStruct.Errorcode))
	}
	return nil
}

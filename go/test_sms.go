package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"team_fighter_go/common"
	"time"
)

func main() {

	// 短信请求 demo

	service_url := "http://122.9.41.215:8000/v1/s/send_msg"
	secret := "66666666"

	sms := Sms{
		AppId:         "gylxt",
		System:        "供应链协同系统",
		PhoneNumbers:  "15922124562",
		SignName:      "供应链协同",
		TemplateCode:  "SMS_225345250",
		TemplateParam: nil,
		OutId:         "YW00001",
		Timestamp:     time.Now().Format("2006-01-02 15:04:05"),
		Sign:          "",
	}
	// 业务参数
	tp := make(map[string]interface{})
	tp["deliveryBh"] = "SYF001002003"
	tp["hospitalName"] = "邵逸夫医院"
	tp_byte, _ := json.Marshal(tp)
	sms.TemplateParam = tp_byte

	// 签名
	sign := Sign(sms.AppId, secret, sms.TemplateCode, sms.Timestamp, string(sms.TemplateParam))
	sms.Sign = sign

	sms_byte, _ := json.Marshal(sms)

	res,err:= http.Post(service_url, "application/json", bytes.NewReader(sms_byte))
	if err != nil {
		fmt.Println(err)
	}
	res_byte,_:= ioutil.ReadAll(res.Body)
	fmt.Println(string(res_byte))
}

type Sms struct {
	AppId         string          `json:"app_id"`
	System        string          `json:"system"`
	PhoneNumbers  string          `json:"phone_numbers"`
	SignName      string          `json:"sign_name"`
	TemplateCode  string          `json:"template_code"`
	TemplateParam json.RawMessage `json:"template_param"`
	OutId         string          `json:"out_id"`
	Timestamp     string          `json:"timestamp"`
	Sign          string          `json:"sign"`
}

//签名生成
func Sign(appId, appSecret, templateCode, timestamp, templateParam string) string {
	var sb bytes.Buffer
	sb.WriteString(appSecret)
	sb.WriteString("app_id")
	sb.WriteString(appId)
	sb.WriteString("template_code")
	sb.WriteString(templateCode)
	sb.WriteString("timestamp")
	sb.WriteString(timestamp)
	sb.WriteString("template_param")
	sb.WriteString(templateParam)
	sb.WriteString(appSecret)
	signString := sb.String()
	signMd5String := common.StringToMd5(signString)
	return strings.ToUpper(signMd5String)
}

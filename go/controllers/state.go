package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"team_fighter_go/common"

	"team_fighter_go/db"
	"time"
)

type StateController struct {
	beego.Controller
}

type httpResponse struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type receiveStateResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
type SendState struct {
	BizID       string `json:"biz_id"`
	ErrCode     string `json:"err_code"`
	ErrMsg      string `json:"err_msg"`
	OutID       string `json:"out_id"`
	PhoneNumber string `json:"phone_number"`
	ReportTime  string `json:"report_time"`
	SendTime    string `json:"send_time"`
	SmsSize     string `json:"sms_size"`
	Success     bool   `json:"success"`
}


type TemplateState struct {
	OrderId         string `json:"order_id"`
	Reason          string `json:"reason"`
	Remark          string `json:"remark"`
	TemplateCode    string `json:"template_code"`
	TemplateContent string `json:"template_content"`
	TemplateName    string `json:"template_name"`
	TemplateStatus  string `json:"template_status"`
	TemplateType    string `json:"template_type"`
	CreateDate      string `json:"create_date"`
}

func (self *StateController) TemplateReceive() {
	defer self.ServeJSON()
	jsonByte := self.Ctx.Input.RequestBody
	logs.Info("Method [template_receive] RequestBody::", string(jsonByte))
	ts_array := make([]TemplateState, 0)
	// JSON 处理
	err_JsonUmarshal := json.Unmarshal(jsonByte, &ts_array)
	if err_JsonUmarshal != nil {
		self.Data["json"] = receiveStateResponse{
			Code: 1,
			Msg:  fmt.Sprintf("解析json发生错误::%s,%s", err_JsonUmarshal, string(jsonByte)),
		}
		logs.Error(fmt.Sprintf("%+v", self.Data["json"]))
		return
	}
	/*
			[{
		        "template_type":"验证码",
		        "reason":"test",
		        "template_name":"短信测试模版1957",
		        "orderId":"14130019",
		        "template_content":"签名测试签名测试签名测试签名测试",
		        "template_status":"approved",
		        "remark":"test",
		        "template_code":"SMS_123123242",
		        "create_date":"2019-05-30 19:58:25"
		    }]
	*/

	// 数据库存储处理
	dbHandler := db.NewHandlerDb()
	// 不使用接口返回的时间，因为该时间就是提交时的时间
	for i, _ := range ts_array {
		ts_array[i].CreateDate = common.FormatDate(time.Now(), common.YYYY_MM_DD_HH_MM_SS)
	}

	_, err_dbExec := dbHandler.Insert(&ts_array)
	if err_dbExec != nil {
		self.Data["json"] = receiveStateResponse{
			Code: 2,
			Msg:  fmt.Sprintf("数据库记录发生错误::%s,%+v", err_dbExec, ts_array),
		}
		logs.Error(fmt.Sprintf("%+v", self.Data["json"]))
		return
	}
	/*
	   重新推送
	   第一次推送失败后，间隔1分钟、5分钟、10分钟、30分钟、60分钟、60分钟、60分钟、60分钟、60分钟后会进行重推，
	   直至推送成功为止。如果推送10次后仍失败，不再重试。
	*/
	self.Data["json"] = receiveStateResponse{
		Code: 0,
		Msg:  "成功",
	}
	return
}

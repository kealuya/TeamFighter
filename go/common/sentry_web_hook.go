package common

//
//import (
//	"github.com/astaxie/beego/logs"
//	"strings"
//	"szht_trip_api/common"
//	"szht_trip_api/msg_send_service"
//)
//
////sentry 收集错误信息的项目集合
//var SENTRY_PROJECT = []string{"travel_jh_vue", "travel_jh_go"}
//
////各项目推送人员id 企业微信openid  zhanbaohua zhangwenjie
//var TRAVEL_JH_VUE_TEAM = []string{"HSDevG1_016","140108199601061959"}
//
////go 项目sentry错误推送人员  zhanbaohua liupengfei
//var TRAVEL_JH_GO_TEAM = []string{"HSDevG1_016","HSDevG1_018"}
//
//type SentryMsg struct {
//	URL         string `json:"url"`         //地址
//	ProjectName string `json:"projectName"` //项目名
//	SentryURL   string `json:"sentryURL"`   //sentry 外网url
//	System      string `json:"system"`      //系统
//	Title       string `json:"title"`       //主题
//}
//
////sentry 错误信息 推送（当前指提供了企业微信推送，有需求可添加短信推送）
//func (pushMsgCtrl *PushMsgController) SentryPost() {
//	resJson := NewJsonStruct(nil)
//
//	defer func() {
//		pushMsgCtrl.Data["json"] = string(common.Marshal(resJson))
//		pushMsgCtrl.ServeJSON()
//	}()
//
//	var sentryMsg = new(SentryMsg)
//	res := pushMsgCtrl.Ctx.Input.RequestBody
//	logs.Debug("sentry调用地址---post 方法" + string(res))
//	common.Unmarshal(res, &sentryMsg)
//	var xmmc = ""
//	for _, v := range SENTRY_PROJECT {
//		if find := strings.Contains(sentryMsg.URL, v); find {
//			xmmc = v
//		}
//	}
//
//	//分项目推送
//	if(xmmc!="" && xmmc =="travel_jh_vue"){
//		for _, v := range TRAVEL_JH_VUE_TEAM {
//			msg_send_service.MsgSendFunc_Szht(v, "【"+xmmc+"】"+"系统出现预期外错误，请及时点击链接处理"+sentryMsg.SentryURL)
//			logs.Info(v+"已下发企业微信推送")
//		}
//	}
//
//	if(xmmc!="" && xmmc =="travel_jh_go"){
//		for _, v := range TRAVEL_JH_GO_TEAM {
//			msg_send_service.MsgSendFunc_Szht(v, "【"+xmmc+"】"+"系统出现预期外错误，请及时点击链接处理"+sentryMsg.SentryURL)
//			logs.Info(v+"已下发企业微信推送")
//		}
//	}
//
//}
//
////大唐订单推送接口
//func (pushMsgCtrl *PushMsgController) SentryGet() {
//	logs.Debug("sentry测试调用地址---get 方法")
//}

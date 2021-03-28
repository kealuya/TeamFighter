package controllers

import "github.com/astaxie/beego"

type ShareController struct {
	beego.Controller
}

func (self *ShareController) TestShare() {
	self.Data["json"] = "分享模块请求成功"
	responseObject := make(map[string]interface{})
	responseObject["username"] = "王鑫钰"
	responseObject["message"] = "测试分享请求"

	self.Data["json"] = httpResponse{
		Success: true,
		Data:    responseObject,
	}
	self.ServeJSON()
}

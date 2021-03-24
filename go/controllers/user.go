package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

func (self *UserController) Test() {

	jsonByte := self.Ctx.Input.RequestBody
	fmt.Println(string(jsonByte))
	self.Data["json"] = "test string display"
	self.ServeJSON()
}

func (self *UserController) Login() {
	defer self.ServeJSON()
	jsonByte := self.Ctx.Input.RequestBody
	logs.Info("Method [Login] RequestBody::", string(jsonByte))
	requestObject := make(map[string]interface{})
	// JSON 处理
	err_JsonUmarshal := json.Unmarshal(jsonByte, &requestObject)
	if err_JsonUmarshal != nil {
		self.Data["json"] = httpResponse{
			Success: false,
			Msg:     fmt.Sprintf("解析json发生错误::%s,%s", err_JsonUmarshal, string(jsonByte)),
		}
		logs.Error(fmt.Sprintf("%+v", self.Data["json"]))
		return
	}

	// logic...

	responseObject := make(map[string]interface{})
	responseObject["result_string"] = "ok"
	responseObject["result_int"] = 123

	self.Data["json"] = httpResponse{
		Success: true,
		Data:    responseObject,
	}
	return
}

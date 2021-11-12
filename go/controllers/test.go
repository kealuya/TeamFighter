package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"team_fighter_go/common"
	"time"
)

type TestController struct {
	beego.Controller
}

var M map[string]interface{}

func init() {
	M = make(map[string]interface{})
}

func (self *TestController) Test() {
	t := time.Now()
	tt := t.Format("15:04:05.000")
	M[tt] = tt
	jsonByte := self.Ctx.Input.RequestBody
	fmt.Println(string(jsonByte))

	fmt.Println(common.GoroutineId())
	go func() {

		fmt.Println("GoroutineId::",common.GoroutineId())

	}()
	time.Sleep(4 * time.Second)
	mJson, _ := json.MarshalIndent(M, "", "\t")
	fmt.Println(string(mJson))
	self.Data["json"] = string(mJson)
	self.ServeJSON()
}

func (self *TestController) TestPost() {
	defer self.ServeJSON()

	fmt.Println(common.GoroutineId())
	jsonByte := self.Ctx.Input.RequestBody
	logs.Info("Method [TestPost] RequestBody::", string(jsonByte))
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

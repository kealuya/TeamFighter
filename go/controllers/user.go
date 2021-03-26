package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go.mongodb.org/mongo-driver/bson"
	"team_fighter_go/common"
	"team_fighter_go/db"
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
	try_err := common.Try(func() {

		requestObject := make(map[string]interface{})
		// JSON 处理
		err_JsonUmarshal := json.Unmarshal(jsonByte, &requestObject)
		common.ErrorHandler(err_JsonUmarshal)
		// logic...
		userid := requestObject["userid"].(string)
		password := requestObject["password"].(string)
		
		collection, ctx := db.ObtainMongoCollection("htjy")
		count, err_countDoc := collection.CountDocuments(ctx, bson.D{{"userid", userid}, {"password", password}})
		common.ErrorHandler(err_countDoc)

		result := make(map[string]interface{})
		result["count"] = count
		self.Data["json"] = httpResponse{
			Success: true,
			Data:    result,
		}
	})
	// 错误处理
	if try_err != nil {
		self.Data["json"] = httpResponse{
			Success: false,
			Msg:     fmt.Sprintf("发生错误::%s", try_err),
		}
		logs.Error(fmt.Sprintf("发生错误::%s - input::%+v", try_err, string(jsonByte)))
	}
}

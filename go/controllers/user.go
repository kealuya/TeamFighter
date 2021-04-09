package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

		returnData := make(map[string]interface{})

		collection, ctx := db.ObtainMongoCollection("htjy")
		collection.FindOne(ctx, bson.D{{"userid", userid}, {"password", password}},
			options.FindOne().SetProjection(
				bson.D{{"userid", 1},
					{"username", 1},
					{"level", 1},
					{"_id", 0},
					{"avatar", 1}})).Decode(&returnData)

		self.Data["json"] = httpResponse{
			Success: true,
			Data:    returnData,
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

func (self *UserController) GetUserAvatar() {
	defer self.ServeJSON()

	jsonByte := self.Ctx.Input.RequestBody
	try_err := common.Try(func() {

		requestObject := make(map[string]interface{})
		// JSON 处理
		err_JsonUmarshal := json.Unmarshal(jsonByte, &requestObject)
		common.ErrorHandler(err_JsonUmarshal)
		// logic...
		userids := requestObject["userids"].([]interface{})

		collection, ctx := db.ObtainMongoCollection("htjy")
		cur, err_FindAvatar := collection.Find(ctx,
			bson.D{{"userid", bson.D{{"$in", userids}}}},
			options.Find().SetProjection(bson.D{{"_id", 0}, {"userid", 1}, {"avatar", 1}}),
		)
		common.ErrorHandler(err_FindAvatar)
		returnData := make(map[string]interface{})
		for cur.Next(ctx) {
			curMap := make(map[string]interface{})
			cur.Decode(&curMap)
			returnData[curMap["userid"].(string)] = curMap["avatar"].(string)
		}

		self.Data["json"] = httpResponse{
			Success: true,
			Data:    returnData,
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

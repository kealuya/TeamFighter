package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"team_fighter_go/common"
	"team_fighter_go/db"
	"time"
)

type TaskController struct {
	beego.Controller
}

func (self *TaskController) GetTaskList() {
	defer self.ServeJSON()

	jsonByte := self.Ctx.Input.RequestBody
	try_err := common.Try(func() {

		requestObject := make(map[string]interface{})
		// JSON 处理
		err_JsonUmarshal := json.Unmarshal(jsonByte, &requestObject)
		common.ErrorHandler(err_JsonUmarshal)
		// logic...
		// 获取userid，query检索条件，及当前页数
		userid := requestObject["userid"].(string)
		query, ok := requestObject["query"].(string)
		if !ok || query == "" {
			query = ""
		}
		page := int(requestObject["page"].(float64))

		//创建返回数据
		returnData := make(map[string]interface{})
		// task总数获取
		collection, ctx := db.ObtainMongoCollection("htjy")
		cur_count, err_Aggregate_count := collection.Aggregate(ctx,
			/*
				1，首先匹配用户
				2，将子doc进行拆分（形成数组）
				3，进行group统计（必须带id），通过{"$sum", 1}，每满足一个条件就加1
				4，重新规整输出的项目仅为 count
			*/
			mongo.Pipeline{
				{
					{"$match",
						bson.D{
							{"userid", userid},
						},
					},
				},
				{
					{"$unwind", "$tasks"},
				},
				{
					{"$group", bson.D{
						{"_id", "$_id"},
						{"total", bson.D{
							{"$sum", 1},
						}},
					}},
				}, {
					{"$project", bson.D{
						{"_id", 0},
						{"count", bson.D{
							{"$sum", "$total"},
						}},
					}},
				},
			})
		if err_Aggregate_count != nil {
			common.ErrorHandler(err_Aggregate_count)
		}
		map_count := make(map[string]int)
		if cur_count.Next(ctx) {
			cur_count.Decode(&map_count)
		}
		returnData["count"] = map_count["count"]
		// task获取
		limit := 10
		skip := (page - 1) * limit

		pipeline := mongo.Pipeline{
			// 匹配userid
			{
				{"$match",
					bson.D{
						{"userid", userid},
					},
				},
			},
			// 对tasks子doc进行拆分，形成数组
			{
				{"$unwind", "$tasks"},
			},

			// 匹配tasks子doc下的todo字段，然后按照正则匹配
			{
				{"$match",
					bson.D{
						{"tasks.todo", bson.D{{"$regex", fmt.Sprintf(`.*%s.*`, query)}}},
					},
				},
			},

			// 按照tasks子doc下的字段进行降序
			{
				{"$sort", bson.D{
					{"tasks.createTime", -1},
				}},
			},
			// 重新规整输出的字段
			{
				{"$project", bson.D{
					{"tasks", 1},
					{"_id", 0},
				}},
			},
			// 分页需要，跳过多少条数据
			{
				{"$skip", skip},
			},
			// 分页需要，一次取前多少条数据
			{
				{"$limit", limit},
			},
		}
		cur, err_Aggregate := collection.Aggregate(ctx, pipeline)
		if err_Aggregate != nil {
			common.ErrorHandler(err_Aggregate)
		}
		curDate := make(map[string]interface{})
		returnTasks := make([]interface{}, 0)

		// 通过游标，获取map类型的数据，并处理成数组
		for cur.Next(ctx) {
			cur.Decode(&curDate)
			returnTasks = append(returnTasks, curDate["tasks"])
		}
		returnData["tasks"] = returnTasks

		// 返回 tasks数组和 count总数
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

func (self *TaskController) GetUserList() {
	defer self.ServeJSON()

	jsonByte := self.Ctx.Input.RequestBody
	try_err := common.Try(func() {

		requestObject := make(map[string]interface{})
		// JSON 处理
		err_JsonUmarshal := json.Unmarshal(jsonByte, &requestObject)
		common.ErrorHandler(err_JsonUmarshal)
		// logic...
		//创建返回数据
		returnData := make(map[string]interface{})
		// task总数获取
		collection, ctx := db.ObtainMongoCollection("htjy")
		cur_FindUserList, err_FindUserList := collection.Find(ctx,
			bson.D{},
			options.Find().SetProjection(bson.D{{"_id", 0},
				{"userid", 1},
				{"username", 1},
				{"avatar", 1},
				{"level", 1},
			}),
		)
		if err_FindUserList != nil {
			common.ErrorHandler(err_FindUserList)
		}
		userList := make([]map[string]interface{}, 0)
		for cur_FindUserList.Next(ctx) {
			userMap := make(map[string]interface{})
			cur_FindUserList.Decode(&userMap)
			userList = append(userList, userMap)
		}
		returnData["userList"] = userList

		// 返回 tasks数组和 count总数
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

func (self *TaskController) SendTaskToUser() {
	defer self.ServeJSON()
	/*
	   //taskNo: "000010",  日期计算吧
	   todo: "测试事项3",   题目
	   todoType: "需求",
	   direction: "in",
	   //progress: NumberInt("75"),
	   fromName: "王鑫钰",
	   fromId: "1587",
	   toName: "任浩",
	   toId: "1209",
	   stars: NumberInt("2"),
	   //state: "wait",
	   info: "测试事项\"2的详情",
	   //createTime: ISODate("2021-04-16T08:47:33.085Z"),
	*/
	jsonByte := self.Ctx.Input.RequestBody
	try_err := common.Try(func() {

		requestObject := make(map[string]interface{})
		// JSON 处理
		err_JsonUmarshal := json.Unmarshal(jsonByte, &requestObject)
		common.ErrorHandler(err_JsonUmarshal)
		// logic...
		// 补充
		// taskNo
		// progress
		// state
		// createTime
		//
		requestObject["taskNo"] = common.FormatDate(time.Now(), common.YYYYMMDDHHMMSS)
		requestObject["progress"] = 0
		requestObject["state"] = "wait"
		requestObject["createTime"] = time.Now()
		direction := "out"
		if requestObject["fromId"] == requestObject["toId"] {
			direction = "none"
		}
		requestObject["direction"] = direction

		bsonM := bson.M{}
		for k, v := range requestObject {
			bsonM[k] = v
		}

		//创建返回数据
		returnData := make(map[string]interface{})
		// task总数获取
		collection, ctx := db.ObtainMongoCollection("htjy")
		// 如果是发送其他人任务
		if bsonM["direction"] == "out" {
			// toUser更新
			bsonM["direction"] = "in"
			_, err_UpdateOne := collection.UpdateOne(ctx,
				bson.D{{"userid", requestObject["toId"]}},
				bson.D{{"$push", bson.D{{"tasks",
					bsonM}}}},
				options.Update())
			common.ErrorHandler(err_UpdateOne)
			bsonM["direction"] = "out"
		}
		// fromUser更新
		ur, err_UpdateOne := collection.UpdateOne(ctx,
			bson.D{{"userid", requestObject["fromId"]}},
			bson.D{{"$push", bson.D{{"tasks",
				bsonM}}}},
			options.Update())
		common.ErrorHandler(err_UpdateOne)

		returnData["modifiedCount"] = ur.ModifiedCount

		// 返回 tasks数组和 count总数
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
func (self *TaskController) UpdateTaskInfo() {
	defer self.ServeJSON()
	/*
	   taskNo: "000010",  日期计算吧
	   todo: "测试事项3",   题目
	   todoType: "需求",
	   direction: "in",
	   progress: NumberInt("75"),
	   fromName: "王鑫钰",
	   fromId: "1587",
	   toName: "任浩",
	   toId: "1209",
	   stars: NumberInt("2"),
	   state: "wait",
	   info: "测试事项\"2的详情",
	   createTime: ISODate("2021-04-16T08:47:33.085Z"),
	*/
	jsonByte := self.Ctx.Input.RequestBody
	try_err := common.Try(func() {

		requestObject := make(map[string]interface{})
		// JSON 处理
		err_JsonUmarshal := json.Unmarshal(jsonByte, &requestObject)
		common.ErrorHandler(err_JsonUmarshal)
		// logic...
		// 补充
		// taskNo
		// progress
		// state
		// createTime
		//
		//taskNo := requestObject["taskNo"]

		//delete(requestObject, "userid")
		bsonM := bson.M{}
		for k, v := range requestObject {
			if !strings.Contains(k, "display") {
				if (k == "createTime" || k == "completeTime") && v != "" {
					t, err_timeParse := time.Parse("2006-01-02 15:04:05", v.(string))
					common.ErrorHandler(err_timeParse, "时间转换错误 %v")
					bsonM[k] = t
				} else {
					bsonM[k] = v
				}
			}
		}
		// 根据用户所选进度，匹配对应任务状态
		if bsonM["progress"].(float64) == 100 {
			bsonM["state"] = "done"
			bsonM["completeTime"] = time.Now()
		} else if bsonM["progress"].(float64) != 0 {
			bsonM["state"] = "confirmed"
		}

		fmt.Println(bsonM)
		//创建返回数据
		// task总数获取
		collection, ctx := db.ObtainMongoCollection("htjy")
		// 如果是发送其他人任务
		if bsonM["direction"] == "out" {
			// toUser更新
			bsonM["direction"] = "in"
			_, err_UpdateOne := collection.UpdateOne(ctx,
				bson.D{{"userid", bsonM["toId"]}, {"tasks.taskNo", bsonM["taskNo"]}},
				bson.D{{"$set", bson.D{{"tasks.$", bsonM}}}},
				options.Update())
			common.ErrorHandler(err_UpdateOne)

			bsonM["direction"] = "out"
			// fromUser更新
			_, err_UpdateOne2 := collection.UpdateOne(ctx,
				bson.D{{"userid", bsonM["fromId"]}, {"tasks.taskNo", bsonM["taskNo"]}},
				bson.D{{"$set", bson.D{{"tasks.$", bsonM}}}},
				options.Update())
			common.ErrorHandler(err_UpdateOne2)
		} else if bsonM["direction"] == "in" {
			// 如果是他人发送自己的任务
			// toUser更新
			_, err_UpdateOne := collection.UpdateOne(ctx,
				bson.D{{"userid", bsonM["toId"]}, {"tasks.taskNo", bsonM["taskNo"]}},
				bson.D{{"$set", bson.D{{"tasks.$", bsonM}}}},
				options.Update())
			common.ErrorHandler(err_UpdateOne)

			// fromUser更新
			bsonM["direction"] = "out"
			_, err_UpdateOne2 := collection.UpdateOne(ctx,
				bson.D{{"userid", bsonM["fromId"]}, {"tasks.taskNo", bsonM["taskNo"]}},
				bson.D{{"$set", bson.D{{"tasks.$", bsonM}}}},
				options.Update())
			common.ErrorHandler(err_UpdateOne2)
		} else {
			// 如果bsonM["direction"] == "none"
			_, err_UpdateOne := collection.UpdateOne(ctx,
				bson.D{{"userid", bsonM["toId"]}, {"tasks.taskNo", bsonM["taskNo"]}},
				bson.D{{"$set", bson.D{{"tasks.$", bsonM}}}},
				options.Update())
			common.ErrorHandler(err_UpdateOne)
		}

		// 返回 tasks数组和 count总数
		self.Data["json"] = httpResponse{
			Success: true,
			Data:    bsonM,
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

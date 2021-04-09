package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"team_fighter_go/common"
	"team_fighter_go/db"
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

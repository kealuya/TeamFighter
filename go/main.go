package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"team_fighter_go/common"
	_ "team_fighter_go/configs" //系统初始化
	"team_fighter_go/controllers"
	"team_fighter_go/db"
	_ "team_fighter_go/routers"
)

func main() {
	//db.ObtainMongoClient()
	//test222()
	beego.ErrorController(&controllers.ErrorController{}) //错误页面处理
	beego.Run()

	//
	//
	//go func() {
	//
	//	beego.ErrorController(&controllers.ErrorController{}) //错误页面处理
	//	beego.Run()
	//
	//}()
	//path := common.GetCurrentPath()
	////path = `C:\Users\Administrator\go\src\szht_sms` //测试环境 path FIXME 正式环境需要注释掉这句话
	//f, err := ioutil.ReadFile(path + "/" + "sms.ico")
	//if err != nil {
	//	f = nil
	//}
	//enc := mahonia.NewEncoder("gb18030")
	//trayhost.EnterLoop(enc.ConvertString("浩天短信sms服务"), f)

}
func mongoTest2() {
	collection, ctx := db.ObtainMongoCollection("htjy")
	count, err_Aggregate_count := collection.Aggregate(ctx, mongo.Pipeline{
		{
			{"$match",
				bson.D{
					{"userid", "1209"},
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
		fmt.Println(err_Aggregate_count)
	}
	returnData := make([]interface{}, 0)
	for count.Next(ctx) {
		curDate := make(map[string]interface{})
		count.Decode(&curDate)
		returnData = append(returnData, curDate)
	}

	fmt.Println(fmt.Sprintf("%+v", returnData))
	fmt.Println(len(returnData))
}

func mongoTest1() {

	//returnData := make(map[string]interface{})

	userid := "1209"
	page := 1
	query := "99"
	collection, ctx := db.ObtainMongoCollection("htjy")
	limit := 3
	skip := (page - 1) * limit

	pipeline := mongo.Pipeline{

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
			{"$match",
				bson.D{
					{"tasks.todo", bson.D{{"$regex", fmt.Sprintf(`.*%s.*`, query)}}},
				},
			},
		},

		{
			{"$sort", bson.D{
				{"tasks.createTime", -1},
			}},
		},
		{
			{"$project", bson.D{
				{"tasks.taskNo", 1},
				{"_id", 0},
			}},
		},

		{
			{"$skip", skip},
		},
		{
			{"$limit", limit},
		},
	}

	cur, err_Aggregate := collection.Aggregate(ctx, pipeline)
	if err_Aggregate != nil {
		common.ErrorHandler(err_Aggregate)
	}
	returnData := make([]interface{}, 0)
	for cur.Next(ctx) {
		curDate := make(map[string]interface{})
		cur.Decode(&curDate)
		returnData = append(returnData, curDate)
	}

	fmt.Println(fmt.Sprintf("%+v", returnData))

	//collection.FindOne(ctx, bson.D{{"userid", "1209"}},
	//	options.FindOne().SetProjection(
	//		bson.D{{"userid", 1},
	//			{"_id", 0},
	//			{"tasks", 1}})).Decode(&returnData)
	//
	//fmt.Println(returnData)
}

func test222() {
	userids := []string{"1209", "1587"}

	collection, ctx := db.ObtainMongoCollection("htjy")

	bson_d := make([]bson.D, 0)
	for _, v := range userids {
		bson_d = append(bson_d, bson.D{
			{"userid", v},
		})

	}
	fmt.Println(bson_d)
	cur, err_FindAvatar := collection.Find(ctx,
		bson.D{{"userid", bson.D{{"$in", []string{"1587", "1208"}}}}},
		options.Find().SetProjection(bson.D{{"_id", 0}, {"userid", 1}, {"avatar", 1}}),
	)
	if err_FindAvatar != nil {
		fmt.Println(err_FindAvatar)
	}

	for cur.Next(ctx) {
		var d interface{}
		cur.Decode(&d)
		fmt.Println(d)
	}
}

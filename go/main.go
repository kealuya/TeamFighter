package main

import (
	"github.com/astaxie/beego"
	_ "team_fighter_go/configs" //系统初始化
	"team_fighter_go/controllers"
	_ "team_fighter_go/routers"
)

func main() {
	//db.ObtainMongoClient()
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

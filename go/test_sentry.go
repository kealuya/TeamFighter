package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"time"
)

func main() {
	/*
		常用参数：
		　　DSN ：项目的地址，用于收集错误信息的 sentry 分配的地址
		　　debug ：是否开启 debug 模式，开启debug，就会把信息打印到控制台上面
		　　release ： 代码的版本号
		　　　　release 版本号，可以确定当前的错误/异常属于哪一个发布的版本
		　　　　可以应用到  sourcemaps 来映射源码
		　　environment : 环境名称
		　　sampleRate : 是否开启随机发送事件给 sentry ，1为100%，0.1 为 10%几率发送事件
		　　attachStacktrace ： 是否开启堆栈跟踪，开启后跟着消息一起收集
		　　beforeSend : 发送前操作
	*/

	/*
		sentry 的api
		　　captureException(exception) : 捕获一个 js 异常，传入一个 exception 对象或者类对象。
		　　captureMessage(message,level) : 捕获一条信息，传入信息内容和信息级别
		　　captureEvent(sentryEvent) : 捕获一个事件，sentryEvent 是手动创建的，自定义的
		　　addBreadcrumb(Breadcrumb) ： 添加一个面包屑，以供接下里的捕获
		　　configureScope((scope)=>{}) : 设置 context 信息到 scope 上面
		　　withScope((scope)=>{}) : 设置一个零时的 scope 信息到 context 上面
	*/
	//配置sentry,自动收集log
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         "http://6ba07683b46a42f4b8ac2a9aa83fbdd6@221.238.155.34:12001/2",
		Environment: "",
		Release:     "syfyy_go@1.0.0",
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	//sentry.ConfigureScope(func(scope *sentry.Scope) {
	//	scope.SetTag("syfyy", "go")
	//	scope.SetContext("device","name")
	//	scope.SetLevel(sentry.LevelWarning)
	//	scope.SetExtra("我的Extra","oookkk")
	//})

	sentry.WithScope(func(scope *sentry.Scope) {

		//scope.SetTag("syfyy", "go")
		//scope.SetUser(sentry.User{
		//	Email:     "kq@qq.com",
		//	ID:        "ttoonnyy",
		//	IPAddress: "",
		//	Username:  "tony",
		//})
		//scope.SetContext("device", "name")
		//scope.SetLevel(sentry.LevelWarning)
		//scope.SetExtra("我的Extra", "oookkk")
		//l := debug.Stack()
		//sentry.CaptureMessage("测试消息6666666666" + string(l))
		//sentry.CaptureException(errors.New("ddddd"))
		e := sentry.NewEvent()
		e.Message = "我是快乐的event6_测试"
		e.Timestamp = time.Now()
		e.Dist = "DistDistDist"//分发中查看
		//e.Fingerprint = []string{"邵逸夫医院_go_记录"} //决定问题是否是一组，会被合并，在【事件】中区分
		e.Logger="LoggerLoggerLogger"//副标题
		m:=make(map[string]string)
		m["Modules1"] ="Modules111"
		m["Modules2"] ="Modules222"
		e.Modules =m
		e.User = sentry.User{
			Email:     "kealuya@126.com",
			ID:        "renhao",
			IPAddress: "192.128.222.111",// 有格式check，需要写正确
			Username:  "2222222",
		}
		e.Level = sentry.LevelError
		extra := make(map[string]interface{})
		// 所有自定义都写到这里
		extra["1"] = "蒙多，想去哪就去哪"
		e.Extra = extra

		sentry.CaptureEvent(e)
	})
	sentry.Flush(time.Second * 5)
	//sentry.AddBreadcrumb(&sentry.Breadcrumb{
	//	Category: "auth",
	//	Message: "Authenticated user " + "test_Breadcrumb",
	//	Level: sentry.LevelDebug,
	//	Data: map[string]interface{}{},
	//
	//});

	fmt.Println("it is ok")
	time.Sleep(5 * time.Second)
}

package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"io/ioutil"
	"reflect"
	"runtime/debug"
	"team_fighter_go/common"
	"time"
)

func main() {

	dd()
	fmt.Println("dd")
}
func dd() {
	extraMap := make(map[string]interface{})
	extraMap["SendToUsers"] = []string{"任浩", "展保华"}
	common.SendInfoToSentry("info消息111111", "测试系统", extraMap)

	defer common.RecoverHandler(nil)

	_, err := ioutil.ReadAll(nil)
	fmt.Println("reflect::", reflect.TypeOf(err).String())

	common.ErrorHandler(err, "%v")
}
func madin() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         "http://de3a21616b704e568fd0924af6ddc8cf@tutou.qcykj.com.cn:9000/2",
		Environment: "",
		Release:     "go@1.0.9",
		Transport:   sentry.NewHTTPSyncTransport(),
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	})
	if err != nil {
		fmt.Println(err)
	}

	main_test()
}

func main_test() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			main_ok(err)
		}
	}()

	fmt.Println("start!!")

	//log.Panicln(errors.New("logic error"))
	var f interface{}
	ff := f.(string)

	fmt.Println("end!!", ff)

}

func main1() {
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
	//err := sentry.Init(sentry.ClientOptions{
	//	Dsn:         "http://de3a21616b704e568fd0924af6ddc8cf@tutou.qcykj.com.cn:9000/2",
	//	Environment: "",
	//	Release:     "go@1.0.9",
	//	// Enable printing of SDK debug messages.
	//	// Useful when getting started or trying to figure something out.
	//	Debug: true,
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//sentry.ConfigureScope(func(scope *sentry.Scope) {
	//	scope.SetTag("syfyy", "go")
	//	scope.SetContext("device","name")
	//	scope.SetLevel(sentry.LevelWarning)
	//	scope.SetExtra("我的Extra","oookkk")
	//})

	//sentry.WithScope(func(scope *sentry.Scope) {

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
	e.Message = "终端测试~！"
	e.Timestamp = time.Now()
	//e.Dist = "分发中查看" //分发中查看
	//e.Fingerprint = []string{"邵逸夫医院_go_记录"} //决定问题是否是一组，会被合并，在【事件】中区分
	e.Logger = "我是快乐的副标题" //副标题

	tag_map := make(map[string]string)
	tag_map["差旅"] = "差旅移动端1.2"
	e.Tags = tag_map
	m := make(map[string]string)
	m["Modules1"] = "Modules111"
	m["Modules2"] = "Modules222"
	e.Modules = m
	e.User = sentry.User{
		Email:     "kealuya@126.com",
		ID:        "renhao",
		IPAddress: "192.128.222.111", // 有格式check，需要写正确
		Username:  "2222222",
	}
	e.Level = sentry.LevelError
	extra := make(map[string]interface{})
	// 附加数据：：所有自定义都写到这里
	extra["1"] = "蒙多，想去哪就去哪"
	extra["error_info"] = string(debug.Stack())
	e.Extra = extra

	bc := make(map[string]interface{})
	bc["step1"] = "step1 detail info"
	bc["step2"] = "step2 detail info"
	bc["step3"] = "step3 detail info"

	// 美国洛杉矶PDT
	loc, _ := time.LoadLocation("America/Los_Angeles")

	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Category:  "auth",
		Message:   "Authenticated user " + "test_Breadcrumb",
		Data:      bc,
		Level:     sentry.LevelDebug,
		Timestamp: time.Now().In(loc),
		Type:      "default",
	})
	e.Exception = []sentry.Exception{
		sentry.Exception{
			Type:       "如果有Exception，那么我就是主标题",
			Value:      "我可是正经的副标题",
			Stacktrace: sentry.ExtractStacktrace(errors.New("st111 error")),
		},
	}

	sentry.CaptureEvent(e)
	//})
	//sentry.Flush(time.Second * 5)

	fmt.Println("it is ok")
	time.Sleep(5 * time.Second)
}

func main_ok(in interface{}) {

	// error错误封装
	var err error
	switch in.(type) {
	case error:
		err = in.(error)
	case string:
		err = errors.New(in.(string))
	}

	fmt.Println("reflect.TypeOf(err).String()::", reflect.TypeOf(err).String())

	fmt.Println("err.Error()::", err.Error())

	e := sentry.NewEvent()

	e.Message = err.Error()
	e.Timestamp = time.Now()
	e.Logger = "我是快乐的副标题" //副标题

	e.User = sentry.User{
		Email:     "kealuya@126.com",
		ID:        "renhao",
		IPAddress: "192.128.222.111", // 有格式check，需要写正确
		Username:  "2222222",
	}
	e.Level = sentry.LevelWarning
	extra := make(map[string]interface{})
	// 附加数据：：所有自定义都写到这里
	extra["1"] = "蒙多，想去哪就去哪"
	extra["Debug Stack"] = string(debug.Stack())
	e.Extra = extra

	bc := make(map[string]interface{})
	bc["step1"] = "step1 detail info"
	bc["step2"] = "step2 detail info"
	bc["step3"] = "step3 detail info"

	// 美国洛杉矶PDT
	loc, _ := time.LoadLocation("America/Los_Angeles")

	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Data:      bc,
		Timestamp: time.Now().In(loc),
	})
	//e.Exception = append(e.Exception, sentry.Exception{
	//	Value:      err.Error()+"11",
	//	Type:       reflect.TypeOf(err).String(),
	//	Stacktrace: sentry.ExtractStacktrace(err),
	//})
	e.Exception = []sentry.Exception{
		sentry.Exception{
			Value:      reflect.TypeOf(err).String(),
			Type:       err.Error(),
			Stacktrace: sentry.NewStacktrace(),
		},
	}

	sentry.CaptureEvent(e)
	//sentry.CaptureException(errors.New("333"))
	//})
	//sentry.Flush(time.Second * 5)

	fmt.Println("it is ok")
	time.Sleep(5 * time.Second)

}

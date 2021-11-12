package common

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"reflect"
	"runtime/debug"
	"time"
)

// fixme 系统版本
var VERSION = "XX系统@1.2.23"

func init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         "http://de3a21616b704e568fd0924af6ddc8cf@tutou.qcykj.com.cn:9000/2",
		Environment: "",
		Release:     VERSION,
		Transport:   sentry.NewHTTPSyncTransport(),
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func SendErrorToSentry(err interface{}, system string, inputParam map[string]interface{}) {
	SendEventToSentry(err, sentry.LevelError, system, inputParam)
}

func SendInfoToSentry(info string, system string, inputParam map[string]interface{}) {
	SendEventToSentry(info, sentry.LevelInfo, system, inputParam)
}

func SendEventToSentry(inputInfo interface{}, errorLevel sentry.Level, system string, extraParam map[string]interface{}) {

	// 对传入对象in进行错误封装（只是为了统一调用方法，Error()））
	var err error
	switch inputInfo.(type) {
	case error:
		err = inputInfo.(error)
	case string:
		err = errors.New(inputInfo.(string))
	}
	// Event做成
	e := sentry.NewEvent()
	// Error()返回errors.New()中的内容
	e.Message = err.Error()
	e.Timestamp = time.Now()
	// 副标题，格式  "系统:模块"
	e.Logger = system
	// 使用用户信息，可以在发生错误时，迅速定位哪个人的操作发生错误
	//e.User = user  暂时无法获得
	e.Level = errorLevel
	extra := make(map[string]interface{})
	// 附加数据：：所有自定义都写到这里
	if extraParam != nil {
		extra = extraParam
	}
	// 默认附带debugStack
	extra["Debug Stack"] = string(debug.Stack())
	e.Extra = extra

	// 除非入参考虑代入ctx，不然step用不到
	//bc := make(map[string]interface{})
	//bc["step1"] = "step1 detail info"
	//bc["step2"] = "step2 detail info"
	//bc["step3"] = "step3 detail info"
	//
	//// 美国洛杉矶PDT
	//loc, _ := time.LoadLocation("America/Los_Angeles")
	//
	//sentry.AddBreadcrumb(&sentry.Breadcrumb{
	//	Data:      bc,
	//	Timestamp: time.Now().In(loc),
	//})

	if errorLevel != sentry.LevelInfo {
		e.Exception = []sentry.Exception{
			sentry.Exception{
				Value:      reflect.TypeOf(err).String(),
				Type:       err.Error(),
				Stacktrace: sentry.NewStacktrace(),
			},
		}
	}

	sentry.CaptureEvent(e)
	//sentry.Flush(time.Second * 5)
}

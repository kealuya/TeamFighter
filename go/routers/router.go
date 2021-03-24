package routers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/goburrow/cache"
	. "team_fighter_go/controllers"
	"time"
)

var cache_md cache.LoadingCache

func init() {

	namespace :=
		beego.NewNamespace("/v1",

			beego.NSNamespace("/b", //base处理
				beego.NSRouter("/login", &UserController{}, "Post:Login"),
			),

			beego.NSNamespace("/t", //test
				beego.NSRouter("/test", &TestController{}, "get:Test"),
				beego.NSRouter("/testPost", &TestController{}, "Post:TestPost"),
			),
			//http://szht.natapp1.cc/v1/r/state_receive
			/*beego.NSNamespace("/r", //receive
				beego.NSRouter("/state_receive", &StateController{}, "post:StateReceive"),
				beego.NSRouter("/template_receive", &StateController{}, "post:TemplateReceive"),
			),*/
		)
	//注册 namespace
	beego.AddNamespace(namespace)

	// Create a new cache
	cache_md = cache.NewLoadingCache(
		func(k cache.Key) (cache.Value, error) {
			return "no", nil
		},
		cache.WithMaximumSize(1000),
		cache.WithRefreshAfterWrite(1*time.Minute),
	)

	//过滤器-传入值记录
	//beego.InsertFilter("/v1/s/send_msg", beego.BeforeExec, filter1)
	//过滤器-传入值记录
	//beego.InsertFilter("/v1/s/add_sms_template", beego.BeforeExec, filter2)

}

func makeResultResponse(s bool, m string) string {
	rr := ResultResponse{
		Success: s,
		Msg:     fmt.Sprintf(m),
		Data:    nil,
	}
	b, _ := json.Marshal(rr)
	return string(b)
}

type IsCache struct{}

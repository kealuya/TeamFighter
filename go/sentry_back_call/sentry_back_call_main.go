package main

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"log"
	"net/http"
)

func readConfig() *ConfigJson {
	fileBytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		logs.Error(err)
		return nil
	}
	jsonConfig := &ConfigJson{}
	err = json.Unmarshal(fileBytes, jsonConfig)
	if err != nil {
		logs.Error(err)
		return nil
	}

	return jsonConfig
}

type ConfigJson struct {
	Info string `json:"info"`
	List []Item `json:"list"`
}
type Item struct {
	SendUsers []SendUsers `json:"send_users"`
	System    string      `json:"system"`
}

type SendUsers struct {
	Name string `json:"name"`
	Qyh  string `json:"qyh"`
	Tel  string `json:"tel"`
}

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		body, _ := ioutil.ReadAll(request.Body)
		var b []byte
		bf := bytes.NewBuffer(b)
		json.Indent(bf, body, "", "    ")
		logs.Info(bf.String())

		si := new(SentryInfo)
		err := json.Unmarshal(body, si)
		if err != nil {
			logs.Error(err)
			writer.WriteHeader(400)
			return
		}

		config := readConfig()
		for _, v := range config.List {
			if v.System == si.Logger {
				SendMsgToDev(v.SendUsers)
			}
		}

		writer.WriteHeader(200)
	})
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		logs.Error(err)
		log.Fatalln(err)
	}

}

func SendMsgToDev() {

}

type SentryInfo struct {
	Culprit string           `json:"culprit,omitempty" example:"team_fighter_go/common in SendEventToSentry"`
	Event   *SentryInfoEvent `json:"event,omitempty"`
	ID      string           `json:"id,omitempty" example:"42"`
	Level   string           `json:"level,omitempty" example:"error"`
	Logger  string           `json:"logger,omitempty" example:"系统"`
	Message string `json:"message,omitempty" example:"come
"`
	Project         string   `json:"project,omitempty" example:"go_test"`
	ProjectName     string   `json:"project_name,omitempty" example:"Go_test"`
	ProjectSlug     string   `json:"project_slug,omitempty" example:"go_test"`
	TriggeringRules []string `json:"triggering_rules,omitempty" example:"Send a notification for new issues"`
	URL             string   `json:"url,omitempty" example:"http://tutou.qcykj.com.cn:9000/sentry/go_test/issues/42/?referrer=webhooks_plugin"`
}

type SentryInfoEvent struct {
	Ref            int                            `json:"_ref,omitempty" example:"2"`
	RefVersion     int                            `json:"_ref_version,omitempty" example:"2"`
	Contexts       *SentryInfoEventContexts       `json:"contexts,omitempty"`
	Culprit        string                         `json:"culprit,omitempty" example:"team_fighter_go/common in SendEventToSentry"`
	EventID        string                         `json:"event_id,omitempty" example:"e577a00a3d454dde8de83701ccc9b888"`
	Exception      *SentryInfoEventException      `json:"exception,omitempty"`
	Extra          *SentryInfoEventExtra          `json:"extra,omitempty"`
	Fingerprint    []string                       `json:"fingerprint,omitempty" example:"{{ default }}"`
	GroupingConfig *SentryInfoEventGroupingConfig `json:"grouping_config,omitempty"`
	Hashes         []string                       `json:"hashes,omitempty" example:"d2692c3dffd865cde8cc9a3ba2843cbe,148eb84e849eeb0fda5d3ad2e79142c6"`
	KeyID          string                         `json:"key_id,omitempty" example:"2"`
	Level          string                         `json:"level,omitempty" example:"error"`
	Location       string                         `json:"location,omitempty" example:"/Users/kealuya/mywork/my_git/TeamFighter/go/common/util_sentry.go"`
	Logentry       *SentryInfoEventLogentry       `json:"logentry,omitempty"`
	Logger         string                         `json:"logger,omitempty" example:"系统"`
	Metadata       *SentryInfoEventMetadata       `json:"metadata,omitempty"`
	Modules        *SentryInfoEventModules        `json:"modules,omitempty"`
	Platform       string                         `json:"platform,omitempty" example:"go"`
	Project        int                            `json:"project,omitempty" example:"2"`
	Received       float64                        `json:"received,omitempty" example:"1.63668562224e+09"`
	Release        string                         `json:"release,omitempty" example:"XX系统@1.2.23"`
	Sdk            *SentryInfoEventSdk            `json:"sdk,omitempty"`
	//Tags           []string                        `json:"tags,omitempty" example:"[level error],[logger 系统],[os.name darwin],[runtime go go1.15.2],[runtime.name go],[sentry:release XX系统@1.2.23],[server_name renhaodeMacBook-Pro.local]"`
	Timestamp float64 `json:"timestamp,omitempty" example:"1.636685621954e+09"`
	//Title          string                         `json:"title,omitempty" example:"come
	//: *errors.fundamental"`
	Type             string `json:"type,omitempty" example:"error"`
	UseRustNormalize bool   `json:"use_rust_normalize,omitempty" example:"true"`
	Version          string `json:"version,omitempty" example:"7"`
}

type SentryInfoEventContexts struct {
	Device  *SentryInfoEventContextsDevice  `json:"device,omitempty"`
	Os      *SentryInfoEventContextsOs      `json:"os,omitempty"`
	Runtime *SentryInfoEventContextsRuntime `json:"runtime,omitempty"`
}

type SentryInfoEventContextsDevice struct {
	Arch   string `json:"arch,omitempty" example:"amd64"`
	NumCPU int    `json:"num_cpu,omitempty" example:"12"`
	Type   string `json:"type,omitempty" example:"device"`
}

type SentryInfoEventContextsOs struct {
	Name string `json:"name,omitempty" example:"darwin"`
	Type string `json:"type,omitempty" example:"os"`
}

type SentryInfoEventContextsRuntime struct {
	GoMaxprocs    int    `json:"go_maxprocs,omitempty" example:"12"`
	GoNumcgocalls int    `json:"go_numcgocalls,omitempty" example:"1"`
	GoNumroutines int    `json:"go_numroutines,omitempty" example:"1"`
	Name          string `json:"name,omitempty" example:"go"`
	Type          string `json:"type,omitempty" example:"runtime"`
	Version       string `json:"version,omitempty" example:"go1.15.2"`
}

type SentryInfoEventException struct {
	Values []SentryInfoEventExceptionValue `json:"values,omitempty"`
}

type SentryInfoEventExceptionValue struct {
	Stacktrace *SentryInfoEventExceptionValueStacktrace `json:"stacktrace,omitempty"`
	Type       string                                   `json:"type,omitempty" example:"come"`
	Value      string                                   `json:"value,omitempty" example:"*errors.fundamental"`
}

type SentryInfoEventExceptionValueStacktrace struct {
	Frames []SentryInfoEventExceptionValueStacktraceFrame `json:"frames,omitempty"`
}

type SentryInfoEventExceptionValueStacktraceFrame struct {
	AbsPath     string   `json:"abs_path,omitempty" example:"/Users/kealuya/mywork/my_git/TeamFighter/go/test_sentry.go"`
	ContextLine string   `json:"context_line,omitempty" example:"	dd()"`
	Filename    string   `json:"filename,omitempty" example:"/Users/kealuya/mywork/my_git/TeamFighter/go/test_sentry.go"`
	Function    string   `json:"function,omitempty" example:"main"`
	InApp       bool     `json:"in_app,omitempty" example:"true"`
	Lineno      int      `json:"lineno,omitempty" example:"16"`
	Module      string   `json:"module,omitempty" example:"main"`
	PostContext []string `json:"post_context,omitempty" example:"	fmt.Println("dd"),},func dd() {,	//extraMap := make(map[string]interface{}),	//extraMap["SendToUsers"] = []string{"任浩", "展保华"}"`
	PreContext  []string `json:"pre_context,omitempty" example:"	"time",),,func main() {,"`
}

type SentryInfoEventExtra struct {
	DebugStack string `json:"Debug Stack,omitempty" example:"goroutine 1 [running]:
runtime/debug.Stack(0xc00010b470, 0x10251a, 0x168ece0)
	/Users/kealuya/go/go1.15.2/src/runtime/debug/stack.go:24 +0x9f
team_fighter_go/common.SendEventToSentry(0x138f240, 0xc000020f20, 0x140b1e8, 0x5, 0x140b69e, 0x6, 0x0)
	/Users/kealuya/mywork/my_git/TeamFighter/go/common/util_sentry.go:64 +0x205
team_fighter_go/common.SendErrorToSentry(...)
	/Users/kealuya/mywork/my_git/TeamFighter/go/common/util_sentry.go:31
team_fighter_go/common.RecoverHandler(0x0)
	/Users/kealuya/mywork/my_git/TeamFighter/go/common/util.go:54 +0xe8
panic(0x138f240, 0xc000020f20)
	/Users/kealuya/go/go1.15.2/src/runtime/panic.go:969 +0x175
log.Panicln(0xc00016be78, 0x1, 0x1)
	/Users/kealuya/go/go1.15.2/src/log/log.go:365 +0xae
team_fighter_go/common.ErrorHandler(0x147d000, 0xc00000e4c0, 0xc000020f00, 0x1, 0x1)
	/Users/kealuya/mywork/my_git/TeamFighter/go/common/util.go:66 +0xeb
main.dd()
	/Users/kealuya/mywork/my_git/TeamFighter/go/test_sentry.go:28 +0xf1
main.main()
	/Users/kealuya/mywork/my_git/TeamFighter/go/test_sentry.go:16 +0x25
"`
}

type SentryInfoEventGroupingConfig struct {
	ID string `json:"id,omitempty" example:"legacy:2019-03-12"`
}

type SentryInfoEventLogentry struct {
	Formatted string `json:"formatted,omitempty" example:"come
"`
}

type SentryInfoEventMetadata struct {
	Filename string `json:"filename,omitempty" example:"/Users/kealuya/mywork/my_git/TeamFighter/go/common/util_sentry.go"`
	Function string `json:"function,omitempty" example:"SendEventToSentry"`
	Type string `json:"type,omitempty" example:"come
"`
	Value string `json:"value,omitempty" example:"*errors.fundamental"`
}

type SentryInfoEventModules struct {
	GithubComAstaxieBeego      string `json:"github.com/astaxie/beego,omitempty" example:"v1.12.1"`
	GithubComGetsentrySentryGo string `json:"github.com/getsentry/sentry-go,omitempty" example:"v0.11.0"`
	GithubComPkgErrors         string `json:"github.com/pkg/errors,omitempty" example:"v0.9.1"`
	GithubComShienaAnsicolor   string `json:"github.com/shiena/ansicolor,omitempty" example:"v0.0.0-20200904210342-c7312218db18"`
	TeamFighterGo              string `json:"team_fighter_go,omitempty" example:"(devel)"`
}

type SentryInfoEventSdk struct {
	Integrations []string                    `json:"integrations,omitempty" example:"ContextifyFrames,Environment,IgnoreErrors,Modules"`
	Name         string                      `json:"name,omitempty" example:"sentry.go"`
	Packages     []SentryInfoEventSdkPackage `json:"packages,omitempty"`
	Version      string                      `json:"version,omitempty" example:"0.11.0"`
}

type SentryInfoEventSdkPackage struct {
	Name    string `json:"name,omitempty" example:"sentry-go"`
	Version string `json:"version,omitempty" example:"0.11.0"`
}

package common

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/exec"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
)

//共通错误recover处理方法 20211111
// fixme
var IS_NEED_SENTRY = true

// input_param 第一个参数是【系统模块名称记述】
// input_param 第二个参数是【传入入参】
func RecoverHandler(f func(recover_err error), input_param ...interface{}) {
	if err := recover(); err != nil {
		// log 打印
		logs.Error(err)
		logs.Error(string(debug.Stack()))
		// sentry发送配置
		if IS_NEED_SENTRY && len(input_param) > 0 {
			i := make(map[string]interface{})
			if len(input_param) == 2 {
				i["入参"] = input_param[1]
			}

			SendErrorToSentry(err, input_param[0].(string), i)
		}
		// 程序错误处理
		if f != nil {
			_, ok := err.(error)
			if ok {
				f(err.(error))
			} else {
				f(errors.New(err.(string)))
			}
		}

	}
}

//共通错误error处理方法 20211111
func ErrorHandler(err error, v ...string) {
	if err != nil {
		if len(v) > 0 {
			log.Panicln(fmt.Errorf(v[0], err))
		} else {
			log.Panicln(err)
		}
	}
}

//获取当前路径
func GetCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

//Float64转String
func Float64ToString(s1 float64, prec int) string {
	return strconv.FormatFloat(s1, 'f', prec, 64) //float64
}

//Float32转String
func Float32ToString(s1 float64, prec int) string {
	//s2 := strconv.FormatFloat(v, 'E', -1, 64)
	return strconv.FormatFloat(s1, 'f', prec, 32) //float32
}

//获取GoroutineID
func GetGoroutineID() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)

	return "goroutine:" + fmt.Sprintf(" %v", n)
}

type WrappedError struct {
	info  error
	input interface{}
}

func (receiver WrappedError) Error() string {
	return receiver.info.Error()
}

//字符串md5加密
func StringToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//字符转base64
func StringToBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

//base64转字符
func Base64ToString(data string) string {
	ds, _ := base64.StdEncoding.DecodeString(data)
	return string(ds)
}

func Try(fn func()) (err error) {
	defer func() {
		if val := recover(); val != nil {
			logs.Error(string(debug.Stack()))
			var ok bool
			err, ok = val.(error)
			if !ok {
				logs.Error("common.Try方法中错误::", val)
				err = errors.New(fmt.Sprintf("%s", val))
			}
		}
	}()

	fn()

	return err
}

/*
   拷贝map
*/
func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}

func GoroutineId() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

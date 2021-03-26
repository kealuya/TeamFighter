package common

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"log"
	"os"
	"os/exec"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
)

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

//共通错误recover处理方法 20200801
func RecoverHandler(f func(rh_err interface{})) {
	if err := recover(); err != nil {
		logs.Error("发生系统错误::", err)
		logs.Error(string(debug.Stack()))
		if f != nil {
			f(err)
		}
	}
}

//共通错误error处理方法 20200801
func ErrorHandler(err error, v ...string) {
	if err != nil {
		if len(v) == 1 {
			log.Panicln(fmt.Errorf(v[0], err))
		} else {
			log.Panicln(err, v)
		}
	}
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
			var ok bool
			err, ok = val.(error)
			if !ok {
				logs.Error("common.Try方法中错误转换错误::", val)
				err = errors.New("common.Try方法中错误转换错误")
			}
		}
	}()

	fn()

	return err
}

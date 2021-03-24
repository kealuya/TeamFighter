package db

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	log_sys "log"
	"sync"
	"xorm.io/xorm"
)

var dbEngine *xorm.Engine
var onceFunc sync.Once

func NewHandlerDb() *xorm.Engine {

	onceFunc.Do(func() {
		engine, err := xorm.NewEngine("mysql", "szht:szht@tcp(114.115.153.238:3306)/sms?charset=utf8")
		if err != nil {
			logs.Error(fmt.Errorf("数据库初始化发生错误::%s", err))
			log_sys.Panicln(err)
		}
		dbEngine = engine
		logs.Info("数据库Engine初始化成功:mysql")
	})
	return dbEngine
}





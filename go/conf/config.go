package conf

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var MyViper *viper.Viper = nil

func init() {
	MyViper = viper.New()
	dir, _ := os.Getwd()
	MyViper.AddConfigPath(dir + "/conf")
	MyViper.SetConfigName("my")
	MyViper.SetConfigType("yaml")
	if err_readInConfig := MyViper.ReadInConfig(); err_readInConfig != nil {
		log.Panicln(err_readInConfig)
	}
}
func GetConfigWithKey(key string) string {
	if MyViper.Get(key) == nil {
		return ""
	}
	return MyViper.Get(key).(string)
}

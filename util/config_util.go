package util

import (
	"fmt"
	"github.com/spf13/viper"
)




var AppConfig *viper.Viper

func LoadConfig() {
	fmt.Println("init config ....")
	AppConfig =viper.New()
	//设置配置文件的名字
	AppConfig.SetConfigName("config")

	//AppConfig.AddConfigPath("/Users/yingfu/go/src/hk_active_report")
	//AppConfig.AddConfigPath("/home/yingfu/go/src/hk_active_report")
	AppConfig.AddConfigPath(BasePath+"/")

	//设置配置文件类型
	AppConfig.SetConfigType("yaml")

	AppConfig.ReadInConfig()
	//AppConfig=V

	//log.Printf("name: %s \n", AppConfig.Get("datasource.url"))
}


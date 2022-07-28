package initialize

import (
	//"github.com/gin-gonic/gin"
	"sao-manage/config"
	"sao-manage/global"
	//"sao-manage/router"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile("setting.dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	//给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 传递给全局变量
	global.Settings = serverConfig
	color.Blue("11111111", global.Settings.LogsAddress)
}

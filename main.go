package main

import (
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"sao-manage/middlewares"
	"time"

	//"go.uber.org/zap"
	"sao-manage/global"
	"sao-manage/initialize"
	//"sao-manage/utils"
)

func main() {
	//1.初始化yaml配置
	initialize.InitConfig()
	initialize.InitLogger()
	//2. 初始化routers
	Router := initialize.Routers()
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	//4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	//5.初始化mysql
	initialize.InitMysqlDB()
	//6. 初始化redis
	initialize.InitRedis()
	// 7. 初始化minIO
	initialize.InitMinIO()
	global.Redis.Set("test","testValue",time.Second)
	value :=global.Redis.Get("test")
	color.Blue(value.Val())
	color.Cyan("go-gin服务开始了")
	//启动gin,并配置端口,global.Settings.Port是yaml配置过的
	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		zap.L().Info("this is hello func", zap.String("error", "启动错误!"))
	}
}
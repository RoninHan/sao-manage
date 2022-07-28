package initialize

import (
	"github.com/gin-gonic/gin"
	"sao-manage/middlewares"

	//"sao-manage/middlewares"
	"sao-manage/router"
)
func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(middlewares.Cors())

	// 路由分组
	ApiGroup := Router.Group("/v1/")

	router.UserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	// 设置跨域中间件
	return Router
}
package router

import (
	"github.com/gin-gonic/gin"
	"sao-manage/controller"
	"sao-manage/middlewares"
)

func BannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banner")
	{
		BannerRouter.POST("list",middlewares.JWTAuth(),middlewares.IsAdminAuth(), controller.GetBannerList)
	}
}

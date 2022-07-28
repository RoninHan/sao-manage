package router
import (
	"github.com/gin-gonic/gin"
	"sao-manage/controller"
	"sao-manage/middlewares"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", controller.PasswordLogin)
		UserRouter.POST("list",middlewares.JWTAuth(),middlewares.IsAdminAuth(), controller.GetUserList)
		UserRouter.POST("putHeaderImage", controller.PutHeaderImage)
		UserRouter.POST("insertUser", controller.InsertUser)
		UserRouter.POST("updatetUserById", controller.UpdatetUser)
	}
}
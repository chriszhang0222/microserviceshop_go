package router

import (
	"github.com/gin-gonic/gin"
	"mxshop/user_web/middleware"
)
import "mxshop/user_web/api"

func InitUserRouter(Router *gin.RouterGroup){
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list",middleware.JWTAUth(), middleware.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PasswordLogin)
		UserRouter.POST("register", api.Register)
	}
}

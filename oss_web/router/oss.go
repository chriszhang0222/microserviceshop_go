package router

import (
	"github.com/gin-gonic/gin"
	"mxshop/oss_web/api"
)

func InitOssRouter(Router *gin.RouterGroup){
	OssRouter := Router.Group("oss")
	{
		//OssRouter.GET("token", middlewares.JWTAuth(), middlewares.IsAdminAuth(), handler.Token)
		OssRouter.GET("/token", api.Token)
		OssRouter.POST("/callback", api.HandlerRequest)
	}
}

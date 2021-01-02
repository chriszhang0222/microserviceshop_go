package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mxshop/oss_web/router"
)

func InitRouters() *gin.Engine{
	Router := gin.Default()
	Router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H {
			"code":http.StatusOK,
			"success":true,
		})
	})
	ApiGroup := Router.Group("/oss")
	router.InitOssRouter(ApiGroup)
	return Router
}

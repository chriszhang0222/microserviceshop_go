package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop/userop_web/middleware"
	"mxshop/userop_web/router"
	"net/http"
)

func InitRouter() *gin.Engine{
	Router := gin.Default()
	Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"success": true,
		})
	})
	Router.Use(middleware.Cors())
	ApiRouter := Router.Group("/v1")
	router.InitAddressRouter(ApiRouter)
	router.InitUserFavRouter(ApiRouter)
	return Router
}

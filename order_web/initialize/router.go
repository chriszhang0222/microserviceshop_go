package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mxshop/order_web/middleware"
	"mxshop/order_web/routers"
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
	ApiGroup := Router.Group("/v1")
	routers.InitShopCart(ApiGroup)
	return Router
}
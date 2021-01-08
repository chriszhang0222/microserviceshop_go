package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop/userop_web/middleware"
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
	return Router
}

package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop/user_web/middleware"
	"net/http"
)
import r "mxshop/user_web/router"

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"success": true,
		})
	})
	Router.Use(middleware.Cors())
	ApiGroup := Router.Group("/v1")
	r.InitUserRouter(ApiGroup)
	r.InitBaseRouter(ApiGroup)
	return Router
}

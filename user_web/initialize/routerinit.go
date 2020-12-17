package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop/user_web/middleware"
)
import r "mxshop/user_web/router"

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Cors())
	ApiGroup := Router.Group("/v1")
	r.InitUserRouter(ApiGroup)
	r.InitBaseRouter(ApiGroup)
	return Router
}

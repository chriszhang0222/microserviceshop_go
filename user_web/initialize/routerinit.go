package initialize

import "github.com/gin-gonic/gin"
import r "mxshop/user_web/router"

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/v1")
	r.InitUserRouter(ApiGroup)
	return Router
}

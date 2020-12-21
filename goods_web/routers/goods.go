package routers

import (
	"github.com/gin-gonic/gin"
	"mxshop/goods_web/api"
	"mxshop/goods_web/middleware"
)

func InitGoodsRouter(Router *gin.RouterGroup){
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("", middleware.JWTAUth(), api.List)
	}
}

package routers

import (
	"github.com/gin-gonic/gin"
	"mxshop/goods_web/api"
)

func InitBrandRouter(Router *gin.RouterGroup){
	BrandRouter := Router.Group("brands")
	{
		BrandRouter.GET("", api.BrandList)
	}
}

package routers

import (
	"github.com/gin-gonic/gin"
	"mxshop/goods_web/api"
)

func InitCategoryRouter(Router *gin.RouterGroup){
	categoryRouter := Router.Group("categories")
	{
		categoryRouter.GET("", api.CategoryList)
	}
}

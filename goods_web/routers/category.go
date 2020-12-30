package routers

import (
	"github.com/gin-gonic/gin"
	"mxshop/goods_web/api"
)

func InitCategoryRouter(Router *gin.RouterGroup){
	categoryRouter := Router.Group("categories")
	{
		categoryRouter.GET("", api.CategoryList)
		categoryRouter.POST("", api.NewCategory)
		categoryRouter.GET("/:id", api.CategoryDetail)
		categoryRouter.PUT("/:id", api.UpdateCategory)
		categoryRouter.DELETE("/:id", api.DeleteCategory)
	}
}

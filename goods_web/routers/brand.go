package routers

import (
	"github.com/gin-gonic/gin"
	"mxshop/goods_web/api"
)

func InitBrandRouter(Router *gin.RouterGroup){
	BrandRouter := Router.Group("brands")
	{
		BrandRouter.GET("", api.BrandList)
		BrandRouter.POST("", api.NewBrand)
		BrandRouter.DELETE("/:id", api.DeleteBrand)
		BrandRouter.PUT("/:id", api.UpdateBrand)

	}
	CategoryBrandRouter := Router.Group("categorybrands")
	{
		CategoryBrandRouter.GET("", api.CategoryBrandList)          // 类别品牌列表页
		CategoryBrandRouter.DELETE("/:id", api.DeleteCategoryBrand) // 删除类别品牌
		CategoryBrandRouter.POST("", api.NewCategoryBrand)       //新建类别品牌
		CategoryBrandRouter.PUT("/:id", api.UpdateCategoryBrand) //修改类别品牌
		CategoryBrandRouter.GET("/:id", api.GetCategoryBrandList) //获取分类的品牌
	}
}

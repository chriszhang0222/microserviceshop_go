package routers

import (
	"github.com/gin-gonic/gin"
	"mxshop/goods_web/api"
)

func InitBannerRouter(Router *gin.RouterGroup){
	BannerRouter := Router.Group("banners")
	{
		BannerRouter.GET("", api.BannerList)
		BannerRouter.POST("", api.NewBanner)
		BannerRouter.PUT("/:id", api.UpdateBanner)
		BannerRouter.DELETE("/:id", api.DeleteBanner)
	}
}

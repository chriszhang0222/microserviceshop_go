package routers

import (
	"github.com/gin-gonic/gin"
	"mxshop/order_web/middleware"
)

func InitShopCart(Router *gin.RouterGroup){
	ShopCartRouter := Router.Group("shopcarts").Use(middleware.JWTAUth())
	{
		ShopCartRouter.GET("")
		ShopCartRouter.POST("")
		ShopCartRouter.DELETE("/:id")
		ShopCartRouter.PATCH("/:id")
	}
}

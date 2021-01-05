package routers

import (
	"github.com/gin-gonic/gin"
	"mxshop/order_web/middleware"
	"mxshop/order_web/api"
)

func InitShopCart(Router *gin.RouterGroup){
	ShopCartRouter := Router.Group("shopcarts").Use(middleware.JWTAUth())
	{
		ShopCartRouter.GET("", api.ShoppingCartList)
		ShopCartRouter.POST("")
		ShopCartRouter.DELETE("/:id")
		ShopCartRouter.PATCH("/:id")
	}
}

package routers

import (
	"github.com/gin-gonic/gin"
	"mxshop/order_web/api"
	"mxshop/order_web/middleware"
)

func InitOrderRouter(Router *gin.RouterGroup){
	OrderRouter := Router.Group("order").Use(middleware.JWTAUth())
	{
		OrderRouter.GET("", api.OrderList)
		OrderRouter.GET("/:id", api.OrderDetail)
		OrderRouter.POST("", api.NewOrder)
	}
}

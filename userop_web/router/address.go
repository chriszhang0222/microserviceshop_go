package router

import (
	"github.com/gin-gonic/gin"
	"mxshop/userop_web/api"
	"mxshop/userop_web/middleware"
)

func InitAddressRouter(Router *gin.RouterGroup){
	AddressRouter := Router.Group("address")
	{
		AddressRouter.GET("", middleware.JWTAUth(), api.AddressList)
		AddressRouter.DELETE("/:id", middleware.JWTAUth(), api.DeleteAddr)
		AddressRouter.POST("", middleware.JWTAUth(), api.NewAddress)
		AddressRouter.PATCH("/:id", middleware.JWTAUth(), api.UpdateAddr)
	}
}
package router

import (
	"github.com/gin-gonic/gin"
	"mxshop/userop_web/api"
	"mxshop/userop_web/middleware"
)

func InitUserFavRouter(Router *gin.RouterGroup){
	UserFavRouter := Router.Group("userfavs").Use(middleware.JWTAUth())
	{
		UserFavRouter.DELETE("/:id", api.DeleteUserFav)
		UserFavRouter.GET("/:id", api.DetailUserFav)
		UserFavRouter.GET("", api.UserFavList)
		UserFavRouter.POST("", api.NewUserFav)
	}
}

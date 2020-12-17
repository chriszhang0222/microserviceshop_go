package middleware

import (
	"github.com/gin-gonic/gin"
	"mxshop/user_web/models"
	"net/http"
)

func IsAdminAuth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUserRole := claims.(*models.CustomClaims).AuthorityId
		if currentUserRole > 1 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "Unauthorized User, not admin",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

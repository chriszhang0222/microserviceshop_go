package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

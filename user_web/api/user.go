package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetUserList(ctx *gin.Context){
	zap.S().Debug("user list")
}

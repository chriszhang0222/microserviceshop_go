package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop/user_web/initialize"
	"mxshop/user_web/global"
)

func main() {

	initialize.InitConfig()
	initialize.InitLogger()
	var port = global.ServerConfig.Port

	Router := initialize.Routers()
	zap.S().Debugf("serve user server at %d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("serve error", err.Error())
	}
}

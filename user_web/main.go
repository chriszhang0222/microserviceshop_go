package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop/user_web/initialize"
)

func main() {

	var port = 8021
	initialize.InitLogger()

	Router := initialize.Routers()
	zap.S().Debugf("serve user server at %d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("serve error", err.Error())
	}
}

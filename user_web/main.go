package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mxshop/user_web/initialize"
	"mxshop/user_web/global"
	"github.com/gin-gonic/gin/binding"
	myvalidator "mxshop/user_web/validators"
)

func main() {

	initialize.InitConfig()
	initialize.InitLogger()
	var port = global.ServerConfig.Port

	if v , ok := binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("mobile", myvalidator.ValidateMobile)
	}
	Router := initialize.Routers()
	zap.S().Debugf("serve user server at %d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("serve error", err.Error())
	}
}

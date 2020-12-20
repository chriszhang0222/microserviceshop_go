package main

import (
	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mxshop/user_web/initialize"
	"mxshop/user_web/global"
	"github.com/gin-gonic/gin/binding"
	myvalidator "mxshop/user_web/validators"
)

var port int

func parsePort(){
	flag.IntVar(&port, "port", global.ServerConfig.Port, "Cache server port")
	flag.Parse()
}
func main() {

	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitSrvConn()
	parsePort()
	if v , ok := binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("mobile", myvalidator.ValidateMobile)
	}
	Router := initialize.Routers()
	zap.S().Debugf("serve user server at %d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("serve error", err.Error())
	}
}

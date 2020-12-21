package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"mxshop/goods_web/initialize"
	"mxshop/goods_web/utils/register"
	"mxshop/goods_web/global"
)
var port int
var registerClient register.RegistryClient

func parsePort(){
	flag.IntVar(&port, "port", global.ServerConfig.Port, "Cache server port")
	flag.Parse()
}

func main(){
	initialize.InitConfig()
	initialize.InitLogger()
	parsePort()
	Router := initialize.InitRouter()
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil{
		zap.S().Panic("serve error", err.Error())
	}
	zap.S().Debugf("serve goods server at %d", port)

}


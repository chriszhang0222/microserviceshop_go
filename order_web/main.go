package order_web

import (
	"flag"
	"go.uber.org/zap"
	"mxshop/order_web/utils/register"
	"mxshop/order_web/global"
	"mxshop/order_web/initialize"
)

var port int
var registerClient register.RegistryClient

func parsePort(){
	flag.IntVar(&port, "port", global.ServerConfig.Port, "Cache server port")
	flag.Parse()
}

func RegisterConsul(addr string, port int, id string, name string, tags ...string){
	consulHost := global.ServerConfig.ConsulInfo.Host
	consulPort := global.ServerConfig.ConsulInfo.Port
	registerClient = register.NewRegistryClient(consulHost, consulPort)
	err := registerClient.Register(addr, port, id, name, tags)
	if err == nil{
		zap.S().Infof("Register to consul %s:%d", consulHost, consulPort)
	}else{
		zap.S().Error("Fail to register consul")
	}
}

func main(){
	initialize.InitConfig()
	initialize.InitLogger()
	parsePort()

}

package main

import (
	"flag"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"mxshop/oss_web/global"
	"mxshop/oss_web/initialize"
	"mxshop/oss_web/utils/register"
	"os"
	"os/signal"
	"syscall"
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
	initialize.InitLogger()
	initialize.InitConfig()
	parsePort()
	Router := initialize.InitRouters()
	serverConfig := global.ServerConfig
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	RegisterConsul(serverConfig.Host, port, serviceId, serverConfig.Name, "mxshop")
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
			zap.S().Panic("serve error", err.Error())
		}
		zap.S().Debugf("serve goods server at %d", port)
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := registerClient.DeRegister(serviceId);err != nil{
		zap.S().Info("Failed to deregister from consul:", err.Error())
	}else{
		zap.S().Info("Deregister from consul")
	}

}

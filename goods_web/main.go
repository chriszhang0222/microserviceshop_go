package main

import (
	"flag"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"mxshop/goods_web/initialize"
	"mxshop/goods_web/utils/register"
	"mxshop/goods_web/global"
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
	initialize.InitConfig()
	initialize.InitLogger()
	parsePort()
	Router := initialize.InitRouter()
	serverConfig := global.ServerConfig
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	RegisterConsul(serverConfig.Host, port, serviceId, serverConfig.Name)
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
		zap.S().Info("Derigister from consul")
	}




}


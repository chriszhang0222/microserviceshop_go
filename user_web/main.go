package main

import (
	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"mxshop/user_web/initialize"
	"mxshop/user_web/global"
	"mxshop/user_web/utils/register"
	"github.com/gin-gonic/gin/binding"
	myvalidator "mxshop/user_web/validators"
	"os"
	"os/signal"
	"syscall"
)

var port int
var register_client register.RegistryClient

func parsePort(){
	flag.IntVar(&port, "port", global.ServerConfig.Port, "Cache server port")
	flag.Parse()
}


func RegisterConsul(addr string, port int, id string, name string, tags ...string){
	consulHost := global.ServerConfig.ConsulInfo.Host
	consulPort := global.ServerConfig.ConsulInfo.Port
	register_client = register.NewRegistryClient(consulHost, consulPort)
	err := register_client.Register(addr, port, id, name, tags)
	if err == nil {
		zap.S().Infof("Register to consul %s:%d", consulHost, consulPort)
	}else{
		zap.S().Error("Fail to register consul")
	}
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
	serverConfig := global.ServerConfig
	serviceId := fmt.Sprintf("%s", uuid.NewV4())

	RegisterConsul(serverConfig.Host, port, serviceId, serverConfig.Name)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("serve error", err.Error())
	}
	zap.S().Debugf("serve user server at %d", port)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := register_client.DeRegister(serviceId); err != nil {
		zap.S().Info("Failed to deregister from consul:", err.Error())
	}else{
		zap.S().Info("Derigister from consul")
	}
}

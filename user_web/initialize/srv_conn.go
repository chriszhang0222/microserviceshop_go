package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop/user_web/global"
	"mxshop/user_web/proto"
)

func InitSrvConn(){
 	cfg := api.DefaultConfig()
 	consulInfo := global.ServerConfig.ConsulInfo
 	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
 	userSrvHost := ""
 	userSrvPort := 0
 	client, err := api.NewClient(cfg)
 	if err != nil{
 		panic(err)
	}
	serviceName := fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvConfig.Name)
	data, err := client.Agent().ServicesWithFilter(serviceName)

	if err != nil{
		panic(err)
	}
	for _, value := range data{
		userSrvHost = value.Address
		userSrvPort = value.Port
	}
	if userSrvHost == ""{
		panic("User service is not available")
	}
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort),  grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("connect to user service failed", "msg", err.Error())
		return
	}
	global.UserSrvClient = proto.NewUserClient(userConn)
	zap.S().Info(fmt.Sprintf("Connected to %s %s:%d", global.ServerConfig.UserSrvConfig.Name, userSrvHost, userSrvPort))
}

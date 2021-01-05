package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop/order_web/global"
	"mxshop/order_web/proto"
)

func InitSrvConn(){
	consulInfo := global.ServerConfig.ConsulInfo
	orderConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.OrderSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))

	if err != nil{
		zap.S().Fatalf("[InitSrvConn] Failed to connect to %s", global.ServerConfig.OrderSrvInfo.Name)
		return
	}
	global.OrderSrvClient = proto.NewOrderClient(orderConn)
}
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

	go func() {
		orderConn, err := grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.OrderSrvInfo.Name),
			grpc.WithInsecure(),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))

		if err != nil{
			zap.S().Fatalf("[InitSrvConn] Failed to connect to %s", global.ServerConfig.OrderSrvInfo.Name)
			return
		}
		global.OrderSrvClient = proto.NewOrderClient(orderConn)
	}()

	go func() {
		goodsConn, err := grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
			grpc.WithInsecure(),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		)
		if err != nil {
			zap.S().Fatal("[InitSrvConn]  Failed to connect to GoodsService")
		}

		global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)
	}()

	go func() {
		invConn, err := grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.InventoryInfo.Name),
			grpc.WithInsecure(),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		)
		if err != nil {
			zap.S().Fatal("[InitSrvConn] Failed to connect to Inventory service")
		}

		global.InventorySrvClient = proto.NewInventoryClient(invConn)
	}()
}
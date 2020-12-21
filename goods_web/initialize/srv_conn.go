package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop/goods_web/global"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"mxshop/goods_web/proto"
)

func InitSrvConn(){
	consulInfo := global.ServerConfig.ConsulInfo
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil{
		zap.S().Fatalf("[InitSrvConn] Failed to connect to %s", global.ServerConfig.GoodsSrvInfo.Name)
		return
	}
	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)
}

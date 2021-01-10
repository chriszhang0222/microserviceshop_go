package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop/userop_web/global"
	"mxshop/userop_web/proto"
)

func InitSrvConn(){
	consulInfo := global.ServerConfig.ConsulInfo
	goodsSrvInfo := global.ServerConfig.GoodsSrvInfo.Name
	userOPSrvInfo := global.ServerConfig.UserOPSrvInfo.Name
	goodsConn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, goodsSrvInfo),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【商品服务失败】")
	}
	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)
	userOpConn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, userOPSrvInfo),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户操作服务失败】")
	}

	global.UserFavClient = proto.NewUserFavClient(userOpConn)
	global.MessageClient = proto.NewMessageClient(userOpConn)
	global.AddressClient = proto.NewAddressClient(userOpConn)
}
package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop/goods_web/global"
	"mxshop/goods_web/proto"
	"mxshop/common/otgrpc"
)

func InitSrvConn(){
	consulInfo := global.ServerConfig.ConsulInfo
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))   //jaeger tracing
	if err != nil{
		zap.S().Fatalf("[InitSrvConn] Failed to connect to %s", global.ServerConfig.GoodsSrvInfo.Name)
		return
	}
	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)
}

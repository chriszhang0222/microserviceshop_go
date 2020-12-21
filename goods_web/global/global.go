package global

import (
	"mxshop/goods_web/proto"
	"mxshop/goods_web/config"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	NacosConfig *config.InitConfig = &config.InitConfig{}
	GoodsSrvClient proto.GoodsClient
)

package global

import (
	"mxshop/order_web/config"
	"mxshop/order_web/proto"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	NacosConfig *config.InitConfig = &config.InitConfig{}
	OrderSrvClient proto.OrderClient
	GoodsSrvClient proto.GoodsClient
	InventorySrvClient proto.InventoryClient
)

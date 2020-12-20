package global

import (
	"mxshop/user_web/config"
	"mxshop/user_web/proto"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	NacosConfig *config.InitConfig = &config.InitConfig{}

	UserSrvClient proto.UserClient

)

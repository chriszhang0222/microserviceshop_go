package global

import (
	"mxshop/user_web/config"
	"mxshop/user_web/proto"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	UserSrvClient proto.UserClient

)

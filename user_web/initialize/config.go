package initialize

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop/user_web/global"
)

func InitConfig(){
	configFileName := "user_web/config.yaml"
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig();err != nil{
		panic(err)
	}
	if err := v.Unmarshal(global.ServerConfig);err != nil{
		panic(err)
	}
	zap.S().Infof("Config &v", global.ServerConfig)

}

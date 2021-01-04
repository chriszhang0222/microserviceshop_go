package initialize

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop/order_web/global"
)

func InitConfig(){
	configFileName := "goods_web/config.yaml"
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil{
		zap.S().Error("Error when read config yaml", err.Error())
		panic(err)
	}
	if err := v.Unmarshal(global.NacosConfig);err != nil {
		panic(err)
	}
	readFromnacos()
	nacosInfo := global.NacosConfig.NacosInfo
	zap.S().Infof("Read config from nacos %s: %d", nacosInfo.Host, nacosInfo.Port)

}

func readFromnacos(){
	nacosInfo := global.NacosConfig.NacosInfo
	sc := []constant.ServerConfig{
		{
			IpAddr: nacosInfo.Host,
			Port: uint64(nacosInfo.Port),
		},
	}
	cc := constant.ClientConfig{
		NamespaceId: nacosInfo.Namespace,
		TimeoutMs: 5000,
		NotLoadCacheAtStart: true,
		LogDir: "tmp/nacos/log",
		CacheDir: "tmp/nacos/cache",
		RotateTime: "1h",
		MaxAge: 3,
		LogLevel: "debug",
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosInfo.DataId,
		Group:  nacosInfo.Group})
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil{
		zap.S().Fatalf("Read nacos config failed： %s", err.Error())
	}

}

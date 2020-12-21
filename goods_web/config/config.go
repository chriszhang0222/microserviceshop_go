package config

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	DataId    string `mapstructure:"dataId" json:"dataId"`
	Group     string `mapstructure:"group" json:"group"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type InitConfig struct {
	NacosInfo NacosConfig `mapstructure:"nacos"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type GoodsSrvConfig struct{
	Host string `mapstructure:"host" json:"host"`
	Port int `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
type ServerConfig struct {
	Host          string        `mapstructure:"host" json:"host"`
	Name          string        `mapstructure:"name" json:"name"`
	Port          int           `mapstructure:"port" json:"port"`
	ConsulInfo    ConsulConfig  `mapstructure:"consul" json:"consul"`
	JWTInfo       JWTConfig     `mapstructure:"jwt" json:"jwt"`
	GoodsSrvInfo  GoodsSrvConfig `mapstructure:"goods_srv" json:"goods_srv"`
}

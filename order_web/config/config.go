package config

import (
	"mxshop/common/config"
)
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

type OrderSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type AlipayConfig struct {
	AppID        string `mapstructure:"app_id" json:"app_id"`
	PrivateKey   string `mapstructure:"private_key" json:"private_key"`
	AliPublicKey string `mapstructure:"ali_public_key" json:"ali_public_key"`
	NotifyURL    string `mapstructure:"notify_url" json:"notify_url"`
	ReturnURL    string `mapstructure:"return_url" json:"return_url"`
}

type ServerConfig struct {
	Host          string        `mapstructure:"host" json:"host"`
	Name          string        `mapstructure:"name" json:"name"`
	Port          int           `mapstructure:"port" json:"port"`
	ConsulInfo    ConsulConfig  `mapstructure:"consul" json:"consul"`
	JWTInfo       JWTConfig     `mapstructure:"jwt" json:"jwt"`
	OrderSrvInfo OrderSrvConfig  `mapstructrue:"order_srv" json:"order_srv"`
	GoodsSrvInfo  OrderSrvConfig  `mapstructure:"goods_srv" json:"goods_srv"`
	InventoryInfo OrderSrvConfig  `mapstructure:"inventory_srv" json:"inventory_srv"`
	AliPayInfo       AlipayConfig   `mapstructure:"alipay" json:"alipay"`
	JaegerInfo config.JaegerConfig  `mapstructure:"jaeger" json:"jaeger"`
}
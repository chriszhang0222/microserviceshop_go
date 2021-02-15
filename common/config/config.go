package config

type JaegerConfig struct {
	ServiceName string `mapstructure:"servicename" json:"servicename"`
	HostPort string `mapstructure:"hostport" json:"hostport"`
}

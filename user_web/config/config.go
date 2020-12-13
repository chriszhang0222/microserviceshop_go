package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name          string        `mapstructure:"name"`
	Port          int           `mapstructure:"port"`
	UserSrvConfig UserSrvConfig `mapstructure:"user_srv"`
}

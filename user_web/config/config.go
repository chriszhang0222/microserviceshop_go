package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string    `mapstructure:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}
type RedisConfig struct{
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Expire int 	`mapstructure:"expire"`
}
type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}
type ServerConfig struct {
	Name          string        `mapstructure:"name"`
	Port          int           `mapstructure:"port"`
	UserSrvConfig UserSrvConfig `mapstructure:"user_srv"`
	JWTInfo       JWTConfig     `mapstructure:"jwt"`
	RedisInfo      RedisConfig   `mapstructure:"redis"`
	ConsulInfo	   ConsulConfig   `mapstructure:"consul"`
}

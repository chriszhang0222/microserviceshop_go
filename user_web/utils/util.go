package utils

import "go.uber.org/zap"

func NewLogger()(*zap.Logger, error){
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./log/myproject.log",
		"stderr",
	}
	return cfg.Build()
}

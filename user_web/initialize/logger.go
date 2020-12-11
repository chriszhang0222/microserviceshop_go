package initialize

import "go.uber.org/zap"

func InitLogger(){
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"././log/myproject.log",
		"stderr",
	}
	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger)
}

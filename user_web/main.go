package main

import (
	"go.uber.org/zap"
)

func NewLogger()(*zap.Logger, error){
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./myproject.log",
		"stderr",
	}
	return cfg.Build()
}
func main(){
	logger, err := NewLogger()
	if err != nil{
		panic(err)
	}
	defer logger.Sync()
	logger.Info("msg...")

}


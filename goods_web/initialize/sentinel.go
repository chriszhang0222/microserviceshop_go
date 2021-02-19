package initialize

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"go.uber.org/zap"
)

func InitSentinel(){
	//TODO: read form nacos
	err := sentinel.InitDefault()
	if err != nil{
		zap.S().Fatalf("unexpected error %v", err)
		return
	}
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource: "goods-srv",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior: flow.Reject,
			Threshold: 10,
			StatIntervalInMs: 5000,
		},
	})
	if err != nil{
		zap.S().Fatalf("unexpected error %v", err)
		return
	}
}

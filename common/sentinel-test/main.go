package main
import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"log"
	//"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
)
func main(){
	err := sentinel.InitDefault()
	if err != nil{
		log.Fatalf("unexpected error %v", err)
	}
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource: "some-test",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior: flow.Reject,
			Threshold: 10,
			StatIntervalInMs: 1000,
		},
	})
	if err != nil{
		log.Fatalf("unexpected error %v", err)
	}
	for i:=0;i<12;i++{
		e,b := sentinel.Entry("some-test", sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			fmt.Println("Rate limited")
		}
		fmt.Println("Pass")
		e.Exit()
	}



}

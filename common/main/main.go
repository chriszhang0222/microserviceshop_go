package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	"regexp"
)

var wg sync.WaitGroup
func startService(ctx context.Context){
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Over")
			return
		default:
			time.Sleep(1*time.Second)
			fmt.Println("Running")
		}
	}
}

func main(){
	email := "zhanr52@mcmaster.ca"
	ok , _ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email)
	fmt.Println(ok)
}

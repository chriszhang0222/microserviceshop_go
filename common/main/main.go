package main

import (
	"context"
	"fmt"
	"sync"
	"time"
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
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go startService(ctx)
	time.Sleep(3*time.Second)
	cancel()
	wg.Wait()
}

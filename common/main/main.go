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



var quit = make(chan bool)
func f(){
	defer wg.Done()
Loop:
	for {
		select {
		case <- quit:
			break Loop
		default:
		}
		fmt.Println("bobby")
		time.Sleep(time.Second)
	}
}

func main(){
}

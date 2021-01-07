package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func produce(c chan <- int){
	for i:=0;i<10;i++{
		fmt.Printf("send %d", i)
		c <- i
	}
}

func get(c <- chan int){
	for i := range c{
		fmt.Printf("get %d", i)
	}
}
func main(){
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}

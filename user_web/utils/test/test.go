package main

import (
	"flag"
	"fmt"
)

func main(){

	var port int
	flag.IntVar(&port, "port", 8001, "Cache server port")
	flag.Parse()
	fmt.Println(port)
}

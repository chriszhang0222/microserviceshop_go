package main

import (
	"fmt"
	"mxshop/oss_web/initialize"
	"mxshop/oss_web/global"
)
func main(){
	initialize.InitConfig()
	fmt.Println(global.Serverconfig.OssInfo)
}

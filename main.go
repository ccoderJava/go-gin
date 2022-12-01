package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"go-gin/pkg/setting"
	"go-gin/routers"
	"log"
	"syscall"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HttpPort)

	//返回一个初始化的endlessServer, 在BeforeBegin时输出当前进程pid.
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err:%v", err)
	}
}

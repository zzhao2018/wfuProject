
package main

import (
	"log"
	"wfuProject/output/generate"
	"wfuProject/output/router"
	"wfuProject/server"
)

func main() {
	err:=server.InitOpt()
	if err!=nil {
		log.Printf("init error,err:%+v\n",err)
		return
	}
	//初始化grpc
	generate.RegisterTestServer(server.GetGrpcServer(),&router.RouterServer{})
	server.Run()
}

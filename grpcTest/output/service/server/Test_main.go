
package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"wfuProject/codeGenerate/output/generate"
	"wfuProject/codeGenerate/output/midware"
	"wfuProject/codeGenerate/output/router"
)

const(
	port=":12306"
)


func main() {
	//监听端口
	lis,err:=net.Listen("tcp",port)
	if err!=nil {
		log.Println("listen tcp error,err:",err)
		return
	}
	defer lis.Close()
	midware.MidWareT()
	grpcServer:=grpc.NewServer()
	//初始化grpc
	generate.RegisterTestServer(grpcServer,&router.RouterServer{})
	err=grpcServer.Serve(lis)
	if err!=nil {
		log.Println("server error,err:",err)
		return
	}
}

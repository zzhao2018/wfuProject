package main

var templeMain=`
package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"wfuProject/codeGenerate/{{.OutputPath}}/generate"
    "wfuProject/codeGenerate/{{.OutputPath}}/router"
)

const(
	port=""
	prome_port=":8080"
)

//监控
func promeGotine(){
	http.Handle("/metrics",promhttp.Handler())
	err:=http.ListenAndServe(prome_port,nil)
	if err!=nil {
		panic(err)
	}
}


func main() {
    //监听
	go promeGotine()
	//处理rpc
	lis,err:=net.Listen("tcp",port)
	if err!=nil {
		log.Println("listen tcp error,err:",err)
		return
	}
	defer lis.Close()
	grpcServer:=grpc.NewServer()
	//初始化grpc
	generate.Register{{.Service.Name}}Server(grpcServer,&router.RouterServer{})
	err=grpcServer.Serve(lis)
	if err!=nil {
		log.Println("server error,err:",err)
		return
	}
}
`

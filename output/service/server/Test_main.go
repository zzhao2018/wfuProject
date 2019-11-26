
package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"wfuProject/output/generate"
	"wfuProject/output/projectUtil"
	"wfuProject/output/router"
)

var(
	port=""
	prome_port=""
)

func init(){
	err:= projectUtil.ParseConfInit(projectUtil.G_TestConfName)
    if err!=nil{
       log.Printf("main ParseConfInit error,err:%+v\n",err)
       return
    }
	port=fmt.Sprintf(":%d", projectUtil.GetParseConfPort())
	prome_port=fmt.Sprintf(":%d", projectUtil.GetParseConfPrometheus().Port)
}

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
	if projectUtil.GetParseConfPrometheus().Switch_on==true {
		go promeGotine()
	}
	//处理rpc
	lis,err:=net.Listen("tcp",port)
	if err!=nil {
		log.Println("listen tcp error,err:",err)
		return
	}
	defer lis.Close()
	grpcServer:=grpc.NewServer()
	//初始化grpc
	generate.RegisterTestServer(grpcServer,&router.RouterServer{})
	err=grpcServer.Serve(lis)
	if err!=nil {
		log.Println("server error,err:",err)
		return
	}
}

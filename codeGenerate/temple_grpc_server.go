package main

var templeData=`
package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"wfuProject/codeGenerate/{{.OutputPath}}/generate"
	"context"
)

const(
	port=""
)


func main() {
	//监听端口
	lis,err:=net.Listen("tcp",port)
	if err!=nil {
		log.Println("listen tcp error,err:",err)
		return
	}
	defer lis.Close()
	grpcServer:=grpc.NewServer()
	//初始化grpc
	generate.Register{{.Service.Name}}Server(grpcServer,&{{.Service.Name}}{})
	err=grpcServer.Serve(lis)
	if err!=nil {
		log.Println("server error,err:",err)
		return
	}
}

type {{.Service.Name}} struct {
}

{{range .Rpc}}
func(s *{{$.Service.Name}}){{.Name}}(ctx context.Context, req *generate.{{.RequestType}})(*generate.{{.ReturnsType}}, error){
   return nil,nil
}
{{end}}
`
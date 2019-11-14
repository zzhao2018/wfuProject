package main

import (
	"awesomeProject/videoLearn/part49/model"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

)

type server struct {
}

func(s *server)Sum(ctx context.Context, req *model.SumRequest) (*model.SumReply, error){
	result:=req.A+req.B
	resp:=&model.SumReply{
		V:                    result,
	}
	return resp,nil
}

func(s *server)Concat(ctx context.Context, req *model.ConcatRequest) (*model.ConcatReply, error){
	result:=req.Data1+req.Data2
	resp:=&model.ConcatReply{
		Result:               result,
	}
	return resp,nil
}

const(
	port=":10244"
)
func main() {
	//监听端口
	lis,err:=net.Listen("tcp",port)
	if err!=nil {
		log.Printf("listen error,err:%+v\n",err)
		return
	}
	defer lis.Close()
	grpcServer:=grpc.NewServer()
	//初始化grpc
	model.RegisterAddServer(grpcServer,&server{})
	err=grpcServer.Serve(lis)
	if err!=nil {
		log.Printf("error,err:%+v\n",err)
	}
}


package controller

import (
	"context"
	"google.golang.org/grpc/metadata"
	"wfuProject/logs"
	"wfuProject/output/generate"
)

type SubController struct {
}

//检查服务参数
func(s *SubController)CheckParams(ctx context.Context, req *generate.SumRequest)(error){
   return nil
}

//方法实现
func(s *SubController)Run(ctx context.Context, req *generate.SumRequest)(*generate.SumReply, error){
	//获得metadata
	md,ok:=metadata.FromIncomingContext(ctx)
	if ok==true {
		logs.Debug(ctx,"get metadata:%+v\n",md)
	}
   return &generate.SumReply{
	   V:                    req.A-req.B,
   },nil
}

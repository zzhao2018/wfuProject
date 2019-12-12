
package controller

import (
	"context"
	"google.golang.org/grpc/metadata"
	"wfuProject/logs"
	"wfuProject/output/generate"
)

type ConcatController struct {
}

//检查服务参数
func(s *ConcatController)CheckParams(ctx context.Context, req *generate.ConcatRequest)(error){
   return nil
}

//方法实现
func(s *ConcatController)Run(ctx context.Context, req *generate.ConcatRequest)(*generate.ConcatReply, error){
	//获得metadata
	md,ok:=metadata.FromIncomingContext(ctx)
	if ok==true {
		logs.Debug(ctx,"get metadata:%+v\n",md)
	}
   return &generate.ConcatReply{
	   Result:               req.Data1+req.Data2,
   },nil
}

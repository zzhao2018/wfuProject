
package controller

import (
	"context"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"time"
	"wfuProject/logs"
	"wfuProject/output/generate"
)

type SumController struct {
}

//检查服务参数
func(s *SumController)CheckParams(ctx context.Context, req *generate.SumRequest)(error){
   return nil
}

//方法实现
func(s *SumController)Run(ctx context.Context, req *generate.SumRequest)(*generate.SumReply, error){
	//获得metadata
	md,ok:=metadata.FromIncomingContext(ctx)
	if ok==true {
		logs.Debug(ctx,"get metadata:%+v\n",md)
	}
	sleepTimeLen:=int64(time.Millisecond)*rand.Int63n(50)
	time.Sleep(time.Duration(sleepTimeLen))
	return &generate.SumReply{
	    V:                    (req.B+req.A),
	},nil
}

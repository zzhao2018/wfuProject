
package controller

import (
	"context"
	"math/rand"
	"time"
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
	sleepTimeLen:=int64(time.Millisecond)*rand.Int63n(50)
	time.Sleep(time.Duration(sleepTimeLen))
	return &generate.SumReply{
	    V:                    (req.B+req.A),
	},nil
}

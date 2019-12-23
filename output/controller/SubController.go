
package controller

import (
	"context"
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
	logs.Debug(ctx,"get data:%+v\n",req)
   return &generate.SumReply{
	   V:                    req.A-req.B,
   },nil
}

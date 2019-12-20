
package controller

import (
	"context"
	"wfuProject/codeGenerate/output/generate"
	"wfuProject/logs"
)

type SumController struct {
}

//检查服务参数
func(s *SumController)CheckParams(ctx context.Context, req *generate.SumRequest)(error){
   return nil
}

//方法实现
func(s *SumController)Run(ctx context.Context, req *generate.SumRequest)(*generate.SumReply, error){
	logs.Debug(ctx,"get data:%+v\n",req)
	return &generate.SumReply{
	    V:                    (req.B+req.A),
	},nil
}

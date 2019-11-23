
package controller

import (
	"wfuProject/codeGenerate/output/generate"
	"context"
)

type SumController struct {
}

//检查服务参数
func(s *SumController)CheckParams(ctx context.Context, req *generate.SumRequest)(error){
   return nil
}

//方法实现
func(s *SumController)Run(ctx context.Context, req *generate.SumRequest)(*generate.SumReply, error){
   return &generate.SumReply{
	   V:                    (req.B+req.A),
   },nil
}

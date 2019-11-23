
package controller

import (
	"wfuProject/codeGenerate/output/generate"
	"context"
)

type SubController struct {
}

//检查服务参数
func(s *SubController)CheckParams(ctx context.Context, req *generate.SumRequest)(error){
   return nil
}

//方法实现
func(s *SubController)Run(ctx context.Context, req *generate.SumRequest)(*generate.SumReply, error){
   return &generate.SumReply{
	   V:                    req.A-req.B,
   },nil
}

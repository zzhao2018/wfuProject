
package controller

import (
	"context"
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
	logs.Debug(ctx,"get data:%+v\n",req)
   return &generate.ConcatReply{
	   Result:               req.Data1+req.Data2,
   },nil
}

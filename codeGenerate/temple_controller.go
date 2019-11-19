package main

var templeData=`
package main

import (
	"wfuProject/codeGenerate/{{.OutputPath}}/generate"
	"context"
)

type {{.Rpc.Name}}Controller struct {
}

//检查服务参数
func(s *{{.Rpc.Name}}Controller)CheckParams(ctx context.Context, req *generate.{{.Rpc.RequestType}})(*generate.{{.Rpc.ReturnsType}}, error){
   return nil,nil
}

//方法实现
func(s *{{.Rpc.Name}}Controller)Run(ctx context.Context, req *generate.{{.Rpc.RequestType}})(*generate.{{.Rpc.ReturnsType}}, error){
   return nil,nil
}
`
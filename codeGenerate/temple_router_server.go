package main

var routerTemp=`
package router

import(
  "context"
  "wfuProject/codeGenerate/{{.OutputPath}}/generate"
)

type RouterServer struct{
}

{{range .Rpc}}
func(r *RouterServer){{.Name}}(ctx context.Context, req *generate.{{.RequestType}})(*generate.{{.ReturnsType}},error){
   return nil,nil
}
{{end}}
`
package main

var routerTemp=`
package router

import(
  "log"
  "context"
  "wfuProject/codeGenerate/{{.OutputPath}}/generate"
  "wfuProject/codeGenerate/{{.OutputPath}}/controller"
)

type RouterServer struct{
}

{{range .Rpc}}
func(r *RouterServer){{.Name}}(ctx context.Context, req *generate.{{.RequestType}})(*generate.{{.ReturnsType}},error){
    {{.Name}}Con:=&controller.{{.Name}}Controller{}
    //检查参数
    err:={{.Name}}Con.CheckParams(ctx,req)
    if err!=nil{
        log.Printf("{{.Name}} CheckParams error,err=%+v\n",err)
        return nil,err
    }
    //操作
    resp,err:={{.Name}}Con.Run(ctx,req)
    if err!=nil{
        log.Printf("{{.Name}} Run error,err=%+v\n",err)
        return nil,err
    }    
    return resp,nil
}
{{end}}
`
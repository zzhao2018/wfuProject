package main

var routerTemp=`
package router

import(
  "fmt"
  "context"
  "wfuProject/codeGenerate/{{.OutputPath}}/generate"
  "wfuProject/codeGenerate/{{.OutputPath}}/controller"
  "wfuProject/midware"
  "wfuProject/server"
  "wfuProject/logs"
)

type RouterServer struct{
}

{{range .Rpc}}
func(r *RouterServer){{.Name}}(ctx context.Context, req *generate.{{.RequestType}})(*generate.{{.ReturnsType}},error){
    ctx=midware.InitServerScanMeta(ctx,"{{$.Service.Name}}","{{.Name}}")
    //初始化traceid
    ctx=logs.SetTraceId(ctx)
    outFunc:=server.BuildUserMidWareChain({{.Name}}MidWare)
    response,err:=outFunc(ctx,req)
    if err!=nil{
        logs.Error(ctx,"{{.Name}} outfunc error,err=%+v\n",err)
        return nil,err
    }
    resp,ok:=response.(*generate.{{.ReturnsType}})
    if ok==false {
        err=fmt.Errorf("{{.Name}}MidWare change type error")
        logs.Error(ctx,"{{.Name}} change type error,err=%+v\n",err)
        return nil,err
    }
    return resp,nil
}

func {{.Name}}MidWare(ctx context.Context, request interface{})(interface{},error){
    req,ok:=request.(*generate.{{.RequestType}})
    if ok==false{
        err:=fmt.Errorf("{{.Name}}MidWare change type error")
        logs.Error(ctx,"{{.Name}}MidWare change type error,err=%+v\n",err)
        return nil,err
    }
    {{.Name}}Con:=&controller.{{.Name}}Controller{}
    //检查参数
    err:={{.Name}}Con.CheckParams(ctx,req)
    if err!=nil{
        logs.Error(ctx,"{{.Name}} CheckParams error,err=%+v\n",err)
        return nil,err
    }
    //操作
    resp,err:={{.Name}}Con.Run(ctx,req)
    if err!=nil{
        logs.Error(ctx,"{{.Name}} Run error,err=%+v\n",err)
        return nil,err
    }    
    return resp,nil
}
{{end}}
`
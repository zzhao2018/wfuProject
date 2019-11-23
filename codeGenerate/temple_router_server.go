package main

var routerTemp=`
package router

import(
  "fmt"
  "log"
  "context"
  "wfuProject/codeGenerate/{{.OutputPath}}/generate"
  "wfuProject/codeGenerate/{{.OutputPath}}/controller"
  "wfuProject/codeGenerate/{{.OutputPath}}/midware"
)

type RouterServer struct{
}

{{range .Rpc}}
func(r *RouterServer){{.Name}}(ctx context.Context, req *generate.{{.RequestType}})(*generate.{{.ReturnsType}},error){
    outFunc:=midware.BuildUserMidWareChain({{.Name}}MidWare)
    response,err:=outFunc(context.Background(),req)
    if err!=nil{
        log.Printf("{{.Name}} outfunc error,err=%+v\n",err)
        return nil,err
    }
    resp,ok:=response.(*generate.{{.ReturnsType}})
    if ok==false {
        err=fmt.Errorf("{{.Name}}MidWare change type error")
        log.Printf("{{.Name}} change type error,err=%+v\n",err)
        return nil,err
    }
    return resp,nil
}

func {{.Name}}MidWare(ctx context.Context, request interface{})(interface{},error){
    req,ok:=request.(*generate.{{.RequestType}})
    if ok==false{
        err:=fmt.Errorf("{{.Name}}MidWare change type error")
        log.Printf("{{.Name}}MidWare change type error,err=%+v\n",err)
        return nil,err
    }
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

package router

import(
  "fmt"
  "context"
  "wfuProject/codeGenerate/output/generate"
  "wfuProject/codeGenerate/output/controller"
  "wfuProject/midware"
  "wfuProject/server"
  "wfuProject/logs"
)

type RouterServer struct{
}


func(r *RouterServer)Sum(ctx context.Context, req *generate.SumRequest)(*generate.SumReply,error){
    ctx=midware.InitServerScanMeta(ctx,"Test","Sum")
    outFunc:=server.BuildUserMidWareChain(SumMidWare)
    response,err:=outFunc(ctx,req)
    if err!=nil{
        logs.Error(ctx,"Sum outfunc error,err=%+v\n",err)
        return nil,err
    }
    resp,ok:=response.(*generate.SumReply)
    if ok==false {
        err=fmt.Errorf("SumMidWare change type error")
        logs.Error(ctx,"Sum change type error,err=%+v\n",err)
        return nil,err
    }
    return resp,nil
}

func SumMidWare(ctx context.Context, request interface{})(interface{},error){
    req,ok:=request.(*generate.SumRequest)
    if ok==false{
        err:=fmt.Errorf("SumMidWare change type error")
        logs.Error(ctx,"SumMidWare change type error,err=%+v\n",err)
        return nil,err
    }
    SumCon:=&controller.SumController{}
    //检查参数
    err:=SumCon.CheckParams(ctx,req)
    if err!=nil{
        logs.Error(ctx,"Sum CheckParams error,err=%+v\n",err)
        return nil,err
    }
    //操作
    resp,err:=SumCon.Run(ctx,req)
    if err!=nil{
        logs.Error(ctx,"Sum Run error,err=%+v\n",err)
        return nil,err
    }    
    return resp,nil
}

func(r *RouterServer)Concat(ctx context.Context, req *generate.ConcatRequest)(*generate.ConcatReply,error){
    ctx=midware.InitServerScanMeta(ctx,"Test","Concat")
    outFunc:=server.BuildUserMidWareChain(ConcatMidWare)
    response,err:=outFunc(ctx,req)
    if err!=nil{
        logs.Error(ctx,"Concat outfunc error,err=%+v\n",err)
        return nil,err
    }
    resp,ok:=response.(*generate.ConcatReply)
    if ok==false {
        err=fmt.Errorf("ConcatMidWare change type error")
        logs.Error(ctx,"Concat change type error,err=%+v\n",err)
        return nil,err
    }
    return resp,nil
}

func ConcatMidWare(ctx context.Context, request interface{})(interface{},error){
    req,ok:=request.(*generate.ConcatRequest)
    if ok==false{
        err:=fmt.Errorf("ConcatMidWare change type error")
        logs.Error(ctx,"ConcatMidWare change type error,err=%+v\n",err)
        return nil,err
    }
    ConcatCon:=&controller.ConcatController{}
    //检查参数
    err:=ConcatCon.CheckParams(ctx,req)
    if err!=nil{
        logs.Error(ctx,"Concat CheckParams error,err=%+v\n",err)
        return nil,err
    }
    //操作
    resp,err:=ConcatCon.Run(ctx,req)
    if err!=nil{
        logs.Error(ctx,"Concat Run error,err=%+v\n",err)
        return nil,err
    }    
    return resp,nil
}

func(r *RouterServer)Sub(ctx context.Context, req *generate.SumRequest)(*generate.SumReply,error){
    ctx=midware.InitServerScanMeta(ctx,"Test","Sub")
    outFunc:=server.BuildUserMidWareChain(SubMidWare)
    response,err:=outFunc(ctx,req)
    if err!=nil{
        logs.Error(ctx,"Sub outfunc error,err=%+v\n",err)
        return nil,err
    }
    resp,ok:=response.(*generate.SumReply)
    if ok==false {
        err=fmt.Errorf("SubMidWare change type error")
        logs.Error(ctx,"Sub change type error,err=%+v\n",err)
        return nil,err
    }
    return resp,nil
}

func SubMidWare(ctx context.Context, request interface{})(interface{},error){
    req,ok:=request.(*generate.SumRequest)
    if ok==false{
        err:=fmt.Errorf("SubMidWare change type error")
        logs.Error(ctx,"SubMidWare change type error,err=%+v\n",err)
        return nil,err
    }
    SubCon:=&controller.SubController{}
    //检查参数
    err:=SubCon.CheckParams(ctx,req)
    if err!=nil{
        logs.Error(ctx,"Sub CheckParams error,err=%+v\n",err)
        return nil,err
    }
    //操作
    resp,err:=SubCon.Run(ctx,req)
    if err!=nil{
        logs.Error(ctx,"Sub Run error,err=%+v\n",err)
        return nil,err
    }    
    return resp,nil
}


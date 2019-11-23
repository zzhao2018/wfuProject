
package router

import(
  "fmt"
  "log"
  "context"
  "wfuProject/codeGenerate/output/generate"
  "wfuProject/codeGenerate/output/controller"
  "wfuProject/codeGenerate/output/midware"
)

type RouterServer struct{
}


func(r *RouterServer)Sum(ctx context.Context, req *generate.SumRequest)(*generate.SumReply,error){
    outFunc:=midware.BuildUserMidWareChain(SumMidWare)
    response,err:=outFunc(context.Background(),req)
    if err!=nil{
        log.Printf("Sum outfunc error,err=%+v\n",err)
        return nil,err
    }
    resp,ok:=response.(*generate.SumReply)
    if ok==false {
        err=fmt.Errorf("SumMidWare change type error")
        log.Printf("Sum change type error,err=%+v\n",err)
        return nil,err
    }
    return resp,nil
}

func SumMidWare(ctx context.Context, request interface{})(interface{},error){
    req,ok:=request.(*generate.SumRequest)
    if ok==false{
        err:=fmt.Errorf("SumMidWare change type error")
        log.Printf("SumMidWare change type error,err=%+v\n",err)
        return nil,err
    }
    SumCon:=&controller.SumController{}
    //检查参数
    err:=SumCon.CheckParams(ctx,req)
    if err!=nil{
        log.Printf("Sum CheckParams error,err=%+v\n",err)
        return nil,err
    }
    //操作
    resp,err:=SumCon.Run(ctx,req)
    if err!=nil{
        log.Printf("Sum Run error,err=%+v\n",err)
        return nil,err
    }    
    return resp,nil
}

func(r *RouterServer)Concat(ctx context.Context, req *generate.ConcatRequest)(*generate.ConcatReply,error){
    outFunc:=midware.BuildUserMidWareChain(ConcatMidWare)
    response,err:=outFunc(context.Background(),req)
    if err!=nil{
        log.Printf("Concat outfunc error,err=%+v\n",err)
        return nil,err
    }
    resp,ok:=response.(*generate.ConcatReply)
    if ok==false {
        err=fmt.Errorf("ConcatMidWare change type error")
        log.Printf("Concat change type error,err=%+v\n",err)
        return nil,err
    }
    return resp,nil
}

func ConcatMidWare(ctx context.Context, request interface{})(interface{},error){
    req,ok:=request.(*generate.ConcatRequest)
    if ok==false{
        err:=fmt.Errorf("ConcatMidWare change type error")
        log.Printf("ConcatMidWare change type error,err=%+v\n",err)
        return nil,err
    }
    ConcatCon:=&controller.ConcatController{}
    //检查参数
    err:=ConcatCon.CheckParams(ctx,req)
    if err!=nil{
        log.Printf("Concat CheckParams error,err=%+v\n",err)
        return nil,err
    }
    //操作
    resp,err:=ConcatCon.Run(ctx,req)
    if err!=nil{
        log.Printf("Concat Run error,err=%+v\n",err)
        return nil,err
    }    
    return resp,nil
}

func(r *RouterServer)Sub(ctx context.Context, req *generate.SumRequest)(*generate.SumReply,error){
    outFunc:=midware.BuildUserMidWareChain(SubMidWare)
    response,err:=outFunc(context.Background(),req)
    if err!=nil{
        log.Printf("Sub outfunc error,err=%+v\n",err)
        return nil,err
    }
    resp,ok:=response.(*generate.SumReply)
    if ok==false {
        err=fmt.Errorf("SubMidWare change type error")
        log.Printf("Sub change type error,err=%+v\n",err)
        return nil,err
    }
    return resp,nil
}

func SubMidWare(ctx context.Context, request interface{})(interface{},error){
    req,ok:=request.(*generate.SumRequest)
    if ok==false{
        err:=fmt.Errorf("SubMidWare change type error")
        log.Printf("SubMidWare change type error,err=%+v\n",err)
        return nil,err
    }
    SubCon:=&controller.SubController{}
    //检查参数
    err:=SubCon.CheckParams(ctx,req)
    if err!=nil{
        log.Printf("Sub CheckParams error,err=%+v\n",err)
        return nil,err
    }
    //操作
    resp,err:=SubCon.Run(ctx,req)
    if err!=nil{
        log.Printf("Sub Run error,err=%+v\n",err)
        return nil,err
    }    
    return resp,nil
}


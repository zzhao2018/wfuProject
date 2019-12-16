
package clientService

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"wfuProject/clientMidware"
	"wfuProject/codeGenerate/output/generate"
)

const(
	addr ="localhost:12306"
)

type OptionFunc func(c *ClientDeal)

type ClientDeal struct {
	serverName string
	traceid string
}

/**************初始化*****************/
func NewClientDeal(serverName string,opts ...OptionFunc)*ClientDeal{
	client:=&ClientDeal{
		serverName:serverName,
	}
	for _,eleFunc:=range opts {
		eleFunc(client)
	}
	return client
}

func OptClientTraceId(traceids string)OptionFunc{
	return func(c *ClientDeal) {
		c.traceid=traceids
	}
}

/*******************方法封装******************/

func(c *ClientDeal)Sum(ctx context.Context, in *generate.SumRequest, opts ...grpc.CallOption) (*generate.SumReply, error) {
	//设置分布式追踪traceid
	ctx=c.contextWithTraceMidWareValue(ctx)
	//启动中间件
	outFunc:=clientMidware.BuildClientMidWareLink(SumClientHandler)
	response,err:=outFunc(ctx,in)
	if err!=nil {
		log.Printf("Sum getResponse error,err:%+v\n",err)
		return nil,err
	}
	responseData,ok:=response.(*generate.SumReply)
	if ok==false {
		err=fmt.Errorf("Sum change response error,err:%+v\n",err)
		log.Printf("Sum change response error\n")
		return nil,err
	}
	return responseData,nil
}

func SumClientHandler(ctx context.Context,request interface{})(interface{},error){
	//连接grpc
	conn,err:=grpc.Dial(addr,grpc.WithInsecure())
	if err!=nil {
		log.Printf("SumClientHandler dial grpc error,err:%+v\n",err)
		return nil,err
	}
	defer conn.Close()
	client:=generate.NewTestClient(conn)
	//获得request
	requestClient,ok:=request.(*generate.SumRequest)
	if ok==false {
		err=fmt.Errorf("change request type error")
		log.Printf("SumClientHandler change type error,err:%+v\n",err)
		return nil,err
	}
	response,err:=client.Sum(ctx,requestClient)
	if err!=nil {
		log.Printf("SumClientHandler Sum error,err:%+v\n",err)
		return nil,err
	}
	return response,nil
}

func(c *ClientDeal)Concat(ctx context.Context, in *generate.ConcatRequest, opts ...grpc.CallOption) (*generate.ConcatReply, error) {
	//设置分布式追踪traceid
	ctx=c.contextWithTraceMidWareValue(ctx)
	//启动中间件
	outFunc:=clientMidware.BuildClientMidWareLink(ConcatClientHandler)
	response,err:=outFunc(ctx,in)
	if err!=nil {
		log.Printf("Concat getResponse error,err:%+v\n",err)
		return nil,err
	}
	responseData,ok:=response.(*generate.ConcatReply)
	if ok==false {
		err=fmt.Errorf("Concat change response error,err:%+v\n",err)
		log.Printf("Concat change response error\n")
		return nil,err
	}
	return responseData,nil
}

func ConcatClientHandler(ctx context.Context,request interface{})(interface{},error){
	//连接grpc
	conn,err:=grpc.Dial(addr,grpc.WithInsecure())
	if err!=nil {
		log.Printf("ConcatClientHandler dial grpc error,err:%+v\n",err)
		return nil,err
	}
	defer conn.Close()
	client:=generate.NewTestClient(conn)
	//获得request
	requestClient,ok:=request.(*generate.ConcatRequest)
	if ok==false {
		err=fmt.Errorf("change request type error")
		log.Printf("ConcatClientHandler change type error,err:%+v\n",err)
		return nil,err
	}
	response,err:=client.Concat(ctx,requestClient)
	if err!=nil {
		log.Printf("ConcatClientHandler Concat error,err:%+v\n",err)
		return nil,err
	}
	return response,nil
}

func(c *ClientDeal)Sub(ctx context.Context, in *generate.SumRequest, opts ...grpc.CallOption) (*generate.SumReply, error) {
	//设置分布式追踪traceid
	ctx=c.contextWithTraceMidWareValue(ctx)
	//启动中间件
	outFunc:=clientMidware.BuildClientMidWareLink(SubClientHandler)
	response,err:=outFunc(ctx,in)
	if err!=nil {
		log.Printf("Sub getResponse error,err:%+v\n",err)
		return nil,err
	}
	responseData,ok:=response.(*generate.SumReply)
	if ok==false {
		err=fmt.Errorf("Sub change response error,err:%+v\n",err)
		log.Printf("Sub change response error\n")
		return nil,err
	}
	return responseData,nil
}

func SubClientHandler(ctx context.Context,request interface{})(interface{},error){
	//连接grpc
	conn,err:=grpc.Dial(addr,grpc.WithInsecure())
	if err!=nil {
		log.Printf("SubClientHandler dial grpc error,err:%+v\n",err)
		return nil,err
	}
	defer conn.Close()
	client:=generate.NewTestClient(conn)
	//获得request
	requestClient,ok:=request.(*generate.SumRequest)
	if ok==false {
		err=fmt.Errorf("change request type error")
		log.Printf("SubClientHandler change type error,err:%+v\n",err)
		return nil,err
	}
	response,err:=client.Sub(ctx,requestClient)
	if err!=nil {
		log.Printf("SubClientHandler Sub error,err:%+v\n",err)
		return nil,err
	}
	return response,nil
}


/********context携带traceid*******/
func(c *ClientDeal) contextWithTraceMidWareValue(ctx context.Context)context.Context{
	ctx=context.WithValue(ctx,clientMidware.ClientTraceServerName{},c.serverName)
	return context.WithValue(ctx,clientMidware.ClientTraceIdKey{},c.traceid)
}


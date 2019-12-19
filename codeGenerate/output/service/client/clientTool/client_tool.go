
package clientTool

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"wfuProject/clientService"
    "wfuProject/clientUtil"
	"wfuProject/codeGenerate/output/generate"
)

const(
	addr ="localhost:12306"
)

type ClientDeal struct {
	clientService *clientService.ClientService
}


/**************初始化*****************/
func NewClientDeal(serverName string,opts ...clientService.OptionFunc)*ClientDeal{
	cs:=clientService.NewClientService(serverName,opts...)
	return &ClientDeal{clientService:cs}
}

/*******************方法封装******************/

func(c *ClientDeal)Sum(ctx context.Context, in *generate.SumRequest, opts ...grpc.CallOption) (*generate.SumReply, error) {
	//调用call
    response,err:=c.clientService.Call(ctx,SumClientHandler,in,"Sum")
	if err!=nil {
		log.Printf("Sum Call error,err:%+v\n",err)
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
	//从短连接中获取conn
	serverMetaData,err:=clientUtil.GetMetaDataFromContext(ctx)
	if err!=nil {
		log.Printf("client_tool SumClientHandler GetMetaDataFromContext error,err:%+v\n",err)
		return nil,err
	}
	conn:=serverMetaData.Conn
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
	//调用call
    response,err:=c.clientService.Call(ctx,ConcatClientHandler,in,"Concat")
	if err!=nil {
		log.Printf("Concat Call error,err:%+v\n",err)
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
	//从短连接中获取conn
	serverMetaData,err:=clientUtil.GetMetaDataFromContext(ctx)
	if err!=nil {
		log.Printf("client_tool SumClientHandler GetMetaDataFromContext error,err:%+v\n",err)
		return nil,err
	}
	conn:=serverMetaData.Conn
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
	//调用call
    response,err:=c.clientService.Call(ctx,SubClientHandler,in,"Sub")
	if err!=nil {
		log.Printf("Sub Call error,err:%+v\n",err)
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
	//从短连接中获取conn
	serverMetaData,err:=clientUtil.GetMetaDataFromContext(ctx)
	if err!=nil {
		log.Printf("client_tool SumClientHandler GetMetaDataFromContext error,err:%+v\n",err)
		return nil,err
	}
	conn:=serverMetaData.Conn
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


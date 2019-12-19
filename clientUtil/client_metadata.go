package clientUtil

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"wfuProject/register"
)

type ClientMetaData struct {
	ServerName string   //服务名称
	NodeList []*register.ServerNode //服务器结点信息
	Traceid string           //分布式追踪id
	SelectNode *register.ServerNode
	Conn *grpc.ClientConn
}

type clientMetaDataKey struct {}

//设置metadata入context
func ContextWithMetaData(ctx context.Context,metadata *ClientMetaData)context.Context{
	return context.WithValue(ctx,clientMetaDataKey{},metadata)
}

//从context中获取数据
func GetMetaDataFromContext(ctx context.Context)(*ClientMetaData,error){
	result,ok:=ctx.Value(clientMetaDataKey{}).(*ClientMetaData)
	if ok==false {
		err:=fmt.Errorf("GetMetaDataFromContext get value error")
		log.Printf("GetMetaDataFromContext get value error\n")
		return nil,err
	}
	return result,nil
}
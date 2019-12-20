package clientMidware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"wfuProject/clientUtil"
	"wfuProject/logs"
)

func NewClientConnectMidWare(nextFunc ClientMidwareFunc) ClientMidwareFunc {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var(
				clientServerMetaData *clientUtil.ClientMetaData
				conn *grpc.ClientConn
			)
			//从connect中获得ip地址
			clientServerMetaData,err=clientUtil.GetMetaDataFromContext(ctx)
			if err!=nil {
				logs.Error(ctx,"midware_connect GetMetaDataFromContext error,err:%+v\n",err)
				return
			}
			//创建连接
			ipAddr:=fmt.Sprintf("%s:%s",clientServerMetaData.SelectNode.Ip,clientServerMetaData.SelectNode.Port)

			fmt.Printf("select ip:%s\n",ipAddr)
			conn,err=grpc.Dial(ipAddr,grpc.WithInsecure())
			if err!=nil {
				logs.Error(ctx,"midware_connect Dial error,err:%+v\n",err)
				return nil,err
			}
			defer conn.Close()
			//传递conn
			clientServerMetaData.Conn=conn
			response,err=nextFunc(ctx,request)
			return
		}
	}

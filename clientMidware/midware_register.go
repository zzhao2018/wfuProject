package clientMidware

import (
	"context"
	"log"
	"wfuProject/clientUtil"
	"wfuProject/register"
)

func NewClientRegisterMidware(registerS register.Register)ClientMidware{
	return func(nextFunc ClientMidwareFunc) ClientMidwareFunc {
		return func(ctx context.Context, request interface{}) (response interface{}, e error) {
			//获取服务名称
			serverMetaData,err:= clientUtil.GetMetaDataFromContext(ctx)
			if err!=nil {
				log.Printf("register midware GetMetaDataFromContext error,err:%+v\n",err)
				return nil,err
			}
			//拉取服务
			serverS,err:=registerS.GetServer(ctx,serverMetaData.ServerName)
			if err!=nil {
				log.Printf("register midware GetServer (%s) error,err:%+v\n",serverMetaData.ServerName,err)
				return nil,err
			}
			//将服务器结点传递
			serverMetaData.NodeList=serverS.Node
			response,err=nextFunc(ctx,request)
			return
		}
	}
}

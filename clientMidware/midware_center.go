package clientMidware

import (
	"context"
)
/***************定义客户端中间件链***************/
type ClientMidwareFunc func(ctx context.Context,request interface{})(interface{},error)

type ClientMidware func(nextFunc ClientMidwareFunc)ClientMidwareFunc

func Chain(begin ClientMidware,otherFunc ...ClientMidware)ClientMidware{
	return func(nextFunc ClientMidwareFunc) ClientMidwareFunc {
		//从后往前连接函数
		for i:=len(otherFunc)-1;i>=0 ;i--  {
			nextFunc=otherFunc[i](nextFunc)
		}
		return begin(nextFunc)
	}
}

/***********************管理中间件**************************/
var clientMidWareLink []ClientMidware=make([]ClientMidware,0)

func AppendClientMidWareLink(clientMidware ClientMidware){
	clientMidWareLink=append(clientMidWareLink,clientMidware)
}


func BuildClientMidWareLink(handler ClientMidwareFunc)ClientMidwareFunc{
	var midClientMidWareLink=make([]ClientMidware,0)
	//增加分布式追踪
	midClientMidWareLink=append(midClientMidWareLink,NewTraceMidware)
	//增加用户链
	if len(clientMidWareLink)>0 {
		midClientMidWareLink=append(midClientMidWareLink,clientMidWareLink...)
	}
	//建立连接
	outFunc:=Chain(midClientMidWareLink[0],midClientMidWareLink[1:]...)
	return outFunc(handler)
}

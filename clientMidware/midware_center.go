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



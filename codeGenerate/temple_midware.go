package main

var midwareTemp=`
package midware

import "context"

//定义func
type MidWareFunc func(ctx context.Context,req interface{})(interface{},error)

//定义midware链
type MidWare func(wareFunc MidWareFunc)MidWareFunc

//自定义中间件
var userMidWareLink []MidWare

//定义链传递函数
/*
  next为下一层调用函数
*/
func Chain(outer MidWare,other ...MidWare)MidWare{
	return func(next MidWareFunc) MidWareFunc {
		//next为传递的最后参数，代表处理函数本身
		//传递链参数
		for i:=len(other)-1;i>=0 ;i--  {
			//设置调用参数
			next=other[i](next)
		}
		//将最前的参数传递给outer
		return outer(next)
	}
}

//添加中间件
func AddUserMidWare(midware ...MidWare){
	userMidWareLink= append(userMidWareLink, midware...)
}

//连接中间件
func BuildUserMidWareChain(handler MidWareFunc)MidWareFunc{
	if len(userMidWareLink)==0 {
		return handler
	}
	headFunc:=Chain(userMidWareLink[0],userMidWareLink[1:]...)
	outerFunc:=headFunc(handler)
	return outerFunc
}
`
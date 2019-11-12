package register

import "context"

type Register interface {
	//获得名字
	Name() string
	//初始化
	Init(ctx context.Context,opts ...RegisterOptionFunc) error
	//注册
	Register(ctx context.Context,server *Server)error
	//取消注册
	UnRegister(ctx context.Context,server *Server)error
	//拉取数据
	GetServer(ctx context.Context,serverName string)(*Server,error)
}
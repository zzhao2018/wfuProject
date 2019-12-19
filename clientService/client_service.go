package clientService

import (
	"context"
	"log"
	"sync"
	"time"
	"wfuProject/clientMidware"
	"wfuProject/clientUtil"
	"wfuProject/register"
	_"wfuProject/register/etcdRegister"
	"wfuProject/requsetBalance"
)

type ClientService struct {
	metadata *clientUtil.ClientMetaData
}

type OptionFunc func(c *ClientService)


//设置全局变量
var(
	globalServiceRegister register.Register
	initServiceRegisterController sync.Once
)

//初始化service
func NewClientService(serviceName string,optFunc ...OptionFunc)*ClientService{
	c:=&ClientService{
		metadata:&clientUtil.ClientMetaData{
			ServerName:serviceName,
		},
	}
	//设置变量
	for _,funcEle:=range optFunc {
		funcEle(c)
	}
	//初始化服务注册
	initServiceRegisterController.Do(func() {
		var err error
		globalServiceRegister,err=register.InitServer(context.TODO(),"etcd",
			register.RegisterInitAddr([]string{"localhost:2379"}),
			register.RegisterInitHeartBeat(5),
			register.RegisterInitTimeOut(5*time.Second),
			)
		if err!=nil {
			log.Printf("client service register init error,err:%+v\n",err)
			return
		}
	})
	return c
}

//初始化追踪id
func OptClientTraceId(traceids string)OptionFunc{
	return func(c *ClientService) {
		c.metadata.Traceid=traceids
	}
}


func(c *ClientService)Call(ctx context.Context,handler clientMidware.ClientMidwareFunc,in interface{},funcName string)(interface{},error){
	ctx= clientUtil.ContextWithMetaData(ctx,c.metadata)
	//启动中间件
	outFunc:=c.BuildClientMidWareLink(handler)
	response,err:=outFunc(ctx,in)
	if err!=nil {
		log.Printf("Call getResponse error,err:%+v\n",err)
		return nil,err
	}
	return response,nil
}



/***********************管理中间件**************************/
func(c *ClientService)BuildClientMidWareLink(handler clientMidware.ClientMidwareFunc)clientMidware.ClientMidwareFunc{
	var midClientMidWareLink=make([]clientMidware.ClientMidware,0)
	//增加分布式追踪
	midClientMidWareLink=append(midClientMidWareLink,clientMidware.NewClientTraceMidware)
	//增加服务发现
	midClientMidWareLink=append(midClientMidWareLink,clientMidware.NewClientRegisterMidware(globalServiceRegister))
	//增加负载均衡
	balance,err:=requsetBalance.GetBalance(requsetBalance.B_RandomWeightBalance)
	if err!=nil {
		log.Printf("client_service BuildClientMidWareLink GetBalance error,err:%+v\n",err)
		return nil
	}
	midClientMidWareLink=append(midClientMidWareLink,clientMidware.NewClientBalanceMidWare(balance))
	//添加短连接
	midClientMidWareLink=append(midClientMidWareLink,clientMidware.NewClientConnectMidWare)
	//建立连接
	outFunc:=clientMidware.Chain(midClientMidWareLink[0],midClientMidWareLink[1:]...)
	return outFunc(handler)
}

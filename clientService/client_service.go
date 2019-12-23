package clientService

import (
	"context"
	"github.com/MXi4oyu/golang.org/x/time/rate"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	zipkin "github.com/uber/jaeger-client-go/transport/zipkin"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
	"wfuProject/clientMidware"
	"wfuProject/clientUtil"
	"wfuProject/logs"
	"wfuProject/midware"
	"wfuProject/register"
	_ "wfuProject/register/etcdRegister"
	"wfuProject/requsetBalance"
)

type ClientService struct {
	metadata *clientUtil.ClientMetaData
	balance requsetBalance.Balance
	limitConf *ClientLimitConf
	limit midware.LimitMid
}

type OptionFunc func(c *ClientService)

//常亮
const(
	logChanSize=100
	logLevel="DEBUG"
)
//设置全局变量
var(
	globalServiceRegister register.Register
	initServiceRegisterController sync.Once
	initServiceLogController sync.Once
	initServicePromeController sync.Once
)

//初始化service
func NewClientService(serviceName string,optFunc ...OptionFunc)*ClientService {
	c := &ClientService{
		metadata: &clientUtil.ClientMetaData{
			ServerName: serviceName,
		},
	}
	//设置变量
	for _, funcEle := range optFunc {
		funcEle(c)
	}
	//初始化服务注册
	initServiceRegisterController.Do(func() {
		var err error
		globalServiceRegister, err = register.InitServer(context.TODO(), "etcd",
			register.RegisterInitAddr([]string{"localhost:2379"}),
			register.RegisterInitHeartBeat(5),
			register.RegisterInitTimeOut(5*time.Second),
		)
		if err != nil {
			log.Printf("client service register init error,err:%+v\n", err)
			return
		}
	})
	//初始化限流
	if c.limitConf != nil {
		c.limit = initLimit(c.limitConf)
	}
	//初始化日志
	initServiceLogController.Do(func() {
		initClientLog(serviceName)
	})
	//初始化监控
	initServicePromeController.Do(func() {
		go initProme()
	})
	//初始化分布式追踪
	trace, _ := initTrace(serviceName)
	opentracing.SetGlobalTracer(trace)
	return c
}

//初始化限流
func OptClientLimit(conf *ClientLimitConf)OptionFunc{
	return func(c *ClientService) {
		c.limitConf=conf
	}
}

//初始化追踪id
func OptClientTraceId(traceids string)OptionFunc{
	return func(c *ClientService) {
		c.metadata.Traceid=traceids
	}
}

//初始化负载均衡类型
func OptClientBalanceType(balanceType int)OptionFunc{
	return func(c *ClientService) {
		if balanceType<0 {
			balanceType=requsetBalance.B_RandomWeightBalance
		}
		balance,err:=requsetBalance.GetBalance(balanceType)
		if err!=nil {
			log.Printf("client_service BuildClientMidWareLink GetBalance error,err:%+v\n",err)
			c.balance,_=requsetBalance.GetBalance(requsetBalance.B_PollingWeightBalance)
			return
		}
		c.balance=balance
	}
}



/*******************初始化配置******************/
//初始化日志
func initClientLog(serviceName string){
	//初始化日志
	logs.InitLog(logChanSize,logs.GetLogLevelFromStr(logLevel),serviceName)
	//初始化控制台日志输出
	outputL:=logs.NewLogConsole()
	logs.LogAddOutPut(outputL)
	//初始化文本日志输出
	fileDir:=`C:\Users\35278\Desktop\测试数据\go_log_test`
	fileName:=`client_test.log`
	logF,err:=logs.NewLogFile(fileDir,fileName)
	if err!=nil {
		log.Printf("initLog NewLogFile error,err:%+v\n",err)
		return
	}
	logs.LogAddOutPut(logF)
}

//初始化监控
func initProme(){
	var prome_port=":8099"
	http.Handle("/metrics",promhttp.Handler())
	err:=http.ListenAndServe(prome_port,nil)
	if err!=nil {
		panic(err)
	}
}

//初始化限流器
type ClientLimitConf struct {
	LimitType string
	Size float64
	TimeDiff int64
	Qbs float64
}

func initLimit(conf *ClientLimitConf)midware.LimitMid{
	var limit midware.LimitMid
	switch conf.LimitType {
	case "counter":
		limit=midware.NewCounterLimit(int64(conf.Size),time.Duration(conf.TimeDiff))
	case "bucket":
		limit=midware.NewBucketLimit(conf.Qbs,conf.Size)
	case "lpLimit":
		limit=midware.NewLPLimit(rate.Limit(conf.Qbs),int(conf.Size))
	default:
		limit=midware.NewLPLimit(rate.Limit(conf.Qbs),int(conf.Size))
	}
	return limit
}

//分布式追踪
func initTrace(serverName string)(opentracing.Tracer,io.Closer){
	//初始化
	transport,err:=zipkin.NewHTTPTransport(
		"",
		zipkin.HTTPBatchSize(1),
		zipkin.HTTPLogger(jaeger.StdLogger),)
	if err!=nil {
		log.Printf("traceInit NewHTTPTransport error:%+v\n",err)
		return nil,nil
	}
	cfg:=&config.Configuration{
		Sampler:             &config.SamplerConfig{
			Type:                     "const",
			Param:                    1,
		},
		Reporter:            &config.ReporterConfig{
			LogSpans:            true,
		},
	}
	r:=jaeger.NewRemoteReporter(transport)
	trace,closer,err:=cfg.New(serverName,config.Reporter(r),config.Logger(jaeger.StdLogger))
	if err!=nil {
		log.Printf("traceInit NewRemoteReporter error:%+v\n",err)
		return nil,nil
	}
	return trace,closer
}

/****************函数调用****************/
func(c *ClientService)Call(ctx context.Context,handler clientMidware.ClientMidwareFunc,in interface{},funcName string)(interface{},error){
	c.metadata.MethodName=funcName
	ctx= clientUtil.ContextWithMetaData(ctx,c.metadata)
	ctx=logs.SetTraceIdFromData(ctx,c.metadata.Traceid)
	//启动中间件
	outFunc:=c.BuildClientMidWareLink(handler)
	response,err:=outFunc(ctx,in)
	if err!=nil {
		logs.Error(ctx,"Call getResponse error,err:%+v\n",err)
		return nil,err
	}
	return response,nil
}



/***********************管理中间件**************************/
func(c *ClientService)BuildClientMidWareLink(handler clientMidware.ClientMidwareFunc)clientMidware.ClientMidwareFunc{
	var midClientMidWareLink=make([]clientMidware.ClientMidware,0)
	//增加监控
	midClientMidWareLink=append(midClientMidWareLink,clientMidware.NewClientPromeMidWare)
	//增加限流
	if c.limit!=nil {
		midClientMidWareLink=append(midClientMidWareLink,clientMidware.NewClientLimitMidWare(c.limit))
	}
	//增加熔断
	midClientMidWareLink=append(midClientMidWareLink,clientMidware.NewClientFuseMidWare)
	//增加分布式追踪
	midClientMidWareLink=append(midClientMidWareLink,clientMidware.NewClientTraceMidware)
	//增加服务发现
	midClientMidWareLink=append(midClientMidWareLink,clientMidware.NewClientRegisterMidware(globalServiceRegister))
	//增加负载均衡
	midClientMidWareLink=append(midClientMidWareLink,clientMidware.NewClientBalanceMidWare(c.balance))
	//添加短连接
	midClientMidWareLink=append(midClientMidWareLink,clientMidware.NewClientConnectMidWare)
	//建立连接
	outFunc:=clientMidware.Chain(midClientMidWareLink[0],midClientMidWareLink[1:]...)
	return outFunc(handler)
}

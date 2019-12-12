package server

import (
	"context"
	"fmt"
	"github.com/MXi4oyu/golang.org/x/time/rate"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/transport/zipkin"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
	"wfuProject/logs"
	"wfuProject/midware"
	_"wfuProject/register/etcdRegister"
	"wfuProject/register"

)

type ServerConf struct {
	grpcServer *grpc.Server
	limit midware.LimitMid
}

var serverConf *ServerConf=&ServerConf{}

func GetGrpcServer()*grpc.Server{
	return serverConf.grpcServer
}

/*********************中间件初始化***********************/
//自定义中间件
var userMidWareLink []midware.MidWare

//添加中间件
func AddUserMidWare(midware ...midware.MidWare){
	userMidWareLink= append(userMidWareLink, midware...)
}

//连接中间件
func BuildUserMidWareChain(handler midware.MidWareFunc)midware.MidWareFunc{
	var midwareLink []midware.MidWare
	//初始化traceid
	midwareLink=append(midwareLink,midware.NewPrepareMidWare())
	//添加监控中间件
	if conf.Prometheus.Switch_on {
		midwareLink=append(midwareLink,midware.PromeScanMidWare)
	}
	//添加分布式追踪
	if conf.Trace.Switch_on {
		midwareLink=append(midwareLink,midware.NewTraceMidWare)
	}
	//添加限流中间件
	if conf.Limit.Switch_on {
		midwareLink=append(midwareLink,midware.NewLimitMidware(serverConf.limit))
	}
	if len(userMidWareLink)!=0 {
		midwareLink=append(midwareLink,userMidWareLink...)
	}
	headFunc:=midware.Chain(midwareLink[0],midwareLink[1:]...)
	outerFunc:=headFunc(handler)
	return outerFunc
}



/**********************模块初始化************************/
func InitOpt()error{
	//解析配置文件
	err:=ParseConfInit(G_TestConfName)
	if err!=nil {
		log.Printf("server InitOpt ParseConfInit error,err:%+v\n",err)
		return err
	}
	log.Printf("conf:%+v\n",conf)
	//初始化grpc
	serverConf.grpcServer=grpc.NewServer()
	//初始化限流器
	if conf.Limit.Switch_on {
		serverConf.limit=initLimit()
	}
	//初始化日志库
	initLog()
	//初始化服务注册
	if conf.Register.Switch_on {
		err=initRegister()
		if err!=nil {
			log.Printf("server InitOpt initRegister error,err:%+v\n",err)
			return err
		}
	}
	//初始化分布式追踪
	if conf.Trace.Switch_on {
		err=initTrace(conf.ServerName)
		if err!=nil {
			log.Printf("server InitOpt initTrace error,err:%+v\n",err)
			return err
		}
	}
	return nil
}


//初始化分布式追踪
func initTrace(serverName string)error{
	//初始化
	transport,err:=zipkin.NewHTTPTransport(
		conf.Trace.Report_addr,
		zipkin.HTTPBatchSize(1),
		zipkin.HTTPLogger(jaeger.StdLogger),
		)
	if err!=nil {
		log.Printf("server initTrace NewHTTPTransport error,err:%+v\n",err)
		return err
	}
	//初始化conf
	cfg:=&config.Configuration{
		Sampler:             &config.SamplerConfig{
			Type:                     conf.Trace.Sample_type,
			Param:                    conf.Trace.Sample_rate,
		},
		Reporter:            &config.ReporterConfig{
			LogSpans:            true,
		},
	}
	//初始化remove
	r:=jaeger.NewRemoteReporter(transport)
	trace,_,err:=cfg.New(serverName,config.Reporter(r),config.Logger(jaeger.StdLogger))
	if err!=nil {
		log.Printf("server initTrace NewCfg error,err:%+v\n",err)
		return err
	}
	//设置全局变量
	opentracing.SetGlobalTracer(trace)
	return nil
}

//初始化服务注册
func initRegister()(error){
	//初始化
	regis,err:=register.InitServer(context.TODO(),"etcd",
		register.RegisterInitTimeOut(conf.Register.TimeOut),
		register.RegisterInitRegisterPath(conf.Register.RegisterPath),
		register.RegisterInitHeartBeat(conf.Register.HeartBeat),
		register.RegisterInitAddr(conf.Register.Addr))
	if err!=nil {
		log.Printf("initRegister InitServer error,err:%+v\n",err)
		return err
	}
	//服务初始化
	serverMeta:=&register.Server{
		Name: conf.ServerName,
		Node: make([]*register.ServerNode,0),
	}
	//获取当前ip
	ips,err:=getLocalIp()
	if err!=nil {
		log.Printf("initRegister getLocalIp error,err:%+v\n",err)
		return err
	}
	serverNode:=&register.ServerNode{
		Ip:     ips,
		Port:   strconv.Itoa(conf.Port),
		Weight: 0,
	}
	serverMeta.Node=append(serverMeta.Node,serverNode)
	//注册服务
	err=regis.Register(context.TODO(),serverMeta)
	if err!=nil {
		log.Printf("initRegister Register error,err:%+v\n",err)
		return err
	}
	return nil
}

//获取当前ip
var ipStore atomic.Value

func getLocalIp()(string,error){
	//查询缓存
	ipstr,ok:=ipStore.Load().(string)
	if ok==true {
		return ipstr,nil
	}
	//无缓存，则读取
	addrs,err:=net.Interfaces()
	if err!=nil {
		log.Printf("getLocalIp InterfaceAddrs error,err:%+v\n",err)
		return "",err
	}
	//获得ip
	for _,addr:=range addrs {
		if (addr.Flags&net.FlagUp)!=0 {
			addrEleArr,err:=addr.Addrs()
			if err!=nil {
				log.Printf("getLocalIp Addrs error,err:%+v\n",err)
				return "",err
			}
			for _,addrEle:=range addrEleArr  {
				if ipnet, ok := addrEle.(*net.IPNet);ok==true&&!ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil{
						ips:=ipnet.IP.String()
						ipStore.Store(ips)
						return ips,nil
					}
				}
			}
		}
	}
	err=fmt.Errorf("getLocalIp ip not found exception")
	return "",err
}


//初始化限流器
func initLimit()midware.LimitMid{
	var limit midware.LimitMid
	switch conf.Limit.Type {
	case "counter":
		limit=midware.NewCounterLimit(int64(conf.Limit.Size),time.Duration(conf.Limit.TimeDiff))
	case "bucket":
		limit=midware.NewBucketLimit(conf.Limit.Qbs,conf.Limit.Size)
	case "lpLimit":
		limit=midware.NewLPLimit(rate.Limit(conf.Limit.Qbs),int(conf.Limit.Size))
	default:
		limit=midware.NewLPLimit(rate.Limit(conf.Limit.Qbs),int(conf.Limit.Size))
	}
	return limit
}

//初始化日志
func initLog(){
	//初始化日志
	logs.InitLog(conf.Logs.ChanSize,logs.GetLogLevelFromStr(conf.Logs.LogLevel),conf.ServerName)
	//初始化控制台日志输出
	outputL:=logs.NewLogConsole()
	logs.LogAddOutPut(outputL)
	//初始化文本日志输出
	fileDir:=`C:\Users\35278\Desktop\测试数据\go_log_test`
	fileName:=`test.log`
	logF,err:=logs.NewLogFile(fileDir,fileName)
	if err!=nil {
		log.Printf("initLog NewLogFile error,err:%+v\n",err)
		return
	}
	logs.LogAddOutPut(logF)

}

/***********************运行执行**********************/
func Run(){
	var port=fmt.Sprintf(":%d", conf.Port)
	//监听
	if conf.Prometheus.Switch_on==true {
		go promeGotine()
	}
	//处理rpc
	lis,err:=net.Listen("tcp",port)
	if err!=nil {
		log.Println("listen tcp error,err:",err)
		return
	}
	if serverConf.grpcServer==nil {
		log.Printf("grpc server is nil\n")
		return
	}
	err=serverConf.grpcServer.Serve(lis)
	if err!=nil {
		log.Printf("server listener error,err:%+v\n",err)
		return
	}
}

//监控
func promeGotine(){
	var prome_port=fmt.Sprintf(":%d", conf.Prometheus.Port)
	http.Handle("/metrics",promhttp.Handler())
	err:=http.ListenAndServe(prome_port,nil)
	if err!=nil {
		panic(err)
	}
}
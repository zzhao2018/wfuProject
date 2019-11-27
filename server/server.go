package server

import (
	"fmt"
	"github.com/MXi4oyu/golang.org/x/time/rate"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
	"wfuProject/midware"
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
	//添加限流中间件
	if conf.Limit.Switch_on {
		midwareLink=append(midwareLink,midware.NewLimitMidware(serverConf.limit))
	}
	//添加监控中间件
	if conf.Prometheus.Switch_on {
		midwareLink=append(midwareLink,midware.PromeScanMidWare)
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
	//初始化
	serverConf.grpcServer=grpc.NewServer()
	serverConf.limit=initLimit()
	return nil
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
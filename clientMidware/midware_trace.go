package clientMidware

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/transport/zipkin"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net/http"
)
type ClientTraceIdKey struct {}
type ClientTraceServerName struct {}

func NewTraceMidware(nextFunc ClientMidwareFunc)ClientMidwareFunc{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//获得追踪id
		traceid,ok:=ctx.Value(ClientTraceIdKey{}).(string)
		if ok==false {
			traceid=""
		}
		//获得服务名称
		serverName,ok:=ctx.Value(ClientTraceServerName{}).(string)
		if ok==false {
			serverName=""
		}
		fmt.Printf("log id:%+v\n",traceid)
		//增加分布式追踪
		if traceid!="" {
			trace,closer:=traceInit(serverName)
			defer closer.Close()
			span:=trace.StartSpan(fmt.Sprintf("span-%s",serverName))
			defer span.Finish()
			req:=injectSpanToCtx(span)
			ctx=metadata.AppendToOutgoingContext(ctx,"wfuproject_trace_label",traceid,"Uber-Trace-Id",req.Header["Uber-Trace-Id"][0])
		}
		//调用处理函数
		response,err=nextFunc(ctx,request)
		return
	}
}

/*******************内部方法***************/
//初始化分布式追踪
func traceInit(serverName string)(opentracing.Tracer,io.Closer){
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

//将span注入http头部
func injectSpanToCtx(span opentracing.Span)*http.Request{
	//将span注册到header中
	req,err:=http.NewRequest("GET","",nil)
	if err!=nil {
		log.Printf("injectSpanToCtx error,err:%+v\n",err)
		return nil
	}
	//注入
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	//设置traceid
	return req
}
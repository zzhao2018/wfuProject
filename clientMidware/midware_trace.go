package clientMidware

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc/metadata"
	"net/http"
	"wfuProject/clientUtil"
	"wfuProject/logs"
	"wfuProject/midware"
)
type ClientTraceIdKey struct {}
type ClientTraceServerName struct {}

func NewClientTraceMidware(nextFunc ClientMidwareFunc)ClientMidwareFunc{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var(
			serverMetaData *clientUtil.ClientMetaData
		)
		//获得追踪id
		serverMetaData,err=clientUtil.GetMetaDataFromContext(ctx)
		if err!=nil {
			logs.Error(ctx,"error,err:%+v\n",err)
			return nil,err
		}
		traceid:=serverMetaData.Traceid
		serverName:=serverMetaData.ServerName
		//增加分布式追踪
		if traceid!="" {
			trace:=opentracing.GlobalTracer()
			//fmt.Sprintf("span-%s-%s",serverName,serverMetaData.MethodName)
			span:=trace.StartSpan(serverName)
			span.SetTag("trace_id",traceid)
			span.LogFields(
				log.String("event","client_trace"),
				log.String("value",traceid),
				log.String("method",serverMetaData.MethodName),
			)
			req:=injectSpanToCtx(ctx,span)
			ctx=metadata.AppendToOutgoingContext(ctx,midware.TraceLabel,traceid,"Uber-Trace-Id",req.Header["Uber-Trace-Id"][0])
			span.Finish()
		}
		//调用处理函数
		response,err=nextFunc(ctx,request)
		return
	}
}

/*******************内部方法***************/
//将span注入http头部
func injectSpanToCtx(ctx context.Context,span opentracing.Span)*http.Request{
	//将span注册到header中
	req,err:=http.NewRequest("GET","",nil)
	if err!=nil {
		logs.Error(ctx,"injectSpanToCtx error,err:%+v\n",err)
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
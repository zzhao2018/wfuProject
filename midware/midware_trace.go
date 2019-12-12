package midware

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc/metadata"
	"strings"
	"wfuProject/logs"
)

func NewTraceMidWare(next MidWareFunc)MidWareFunc{
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		//从ctx中获得traceid
		md,ok:=metadata.FromIncomingContext(ctx)
		if ok==false {
			md=metadata.Pairs()
		}
		//从ctx中解析span
		trace:=opentracing.GlobalTracer()
		spanCtx,err:=trace.Extract(opentracing.HTTPHeaders,TraceHeadersCarrier(md))
		if err!=nil {
			logs.Error(ctx,"NewMidwareTrace Extract error,err:%+v\n",err)
			return
		}
		//构造span标志
		serverMetaData:=GetServerScanMeta(ctx)
		span:=trace.StartSpan(serverMetaData.ServerName,ext.RPCServerOption(spanCtx))
		//标志traceid
		trace_id,err:=logs.GetTraceId(ctx)
		if err!=nil {
			trace_id=""
		}
		span.SetTag("trace_id",trace_id)
		resp,err=next(ctx,req)
		if err!=nil {
			logs.Error(ctx,"NewMidwareTrace SetTag error,err:%+v\n",err)
			span.LogFields(
				log.String("event","error"),
				log.String("value",err.Error()),
				)
		}
		span.Finish()
		return
	}
}


//重构md数据结构
type TraceHeadersCarrier metadata.MD

// Set conforms to the TextMapWriter interface.
func (c TraceHeadersCarrier) Set(key, val string) {
	k := strings.ToLower(key)
	c[k]=append(c[k],val)
}

// ForeachKey conforms to the TextMapReader interface.
func (c TraceHeadersCarrier) ForeachKey(handler func(key, val string) error) error {
	for k, vals := range c {
		for _, v := range vals {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}
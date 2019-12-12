package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/transport/zipkin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
	"wfuProject/output/generate"
)

//初始化分布式追踪
func Init(serverName string)(opentracing.Tracer,io.Closer){
	//初始化
	transport,err:=zipkin.NewHTTPTransport(
		"",
		zipkin.HTTPBatchSize(1),
		zipkin.HTTPLogger(jaeger.StdLogger),)
	if err!=nil {
		panic(err.Error())
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
		panic(err.Error())
	}
	return trace,closer
}


func main() {
	conn,err:=grpc.Dial("localhost:12306",grpc.WithInsecure())
	if err!=nil {
		log.Printf("dial error,err:%+v\n",err)
		return
	}

	/************分布式追踪**************/
	//添加span
	trace,closer:=Init("test-wfuProject")
	defer closer.Close()
	span:=trace.StartSpan("test-span")
	defer span.Finish()
	//将span注册到header中
	req:=injectSpanToCtx(span)
	//设置traceid
	ctx:=metadata.AppendToOutgoingContext(context.Background(),"wfuproject_trace_label","12345","Uber-Trace-Id",req.Header["Uber-Trace-Id"][0])

	client:= generate.NewTestClient(conn)
	reply,err:=client.Concat(ctx,&generate.ConcatRequest{
		Data1:                "hello,",
		Data2:                "zzh",
	})
	if err!=nil {
		log.Printf("Concat error,err:%+v\n",err)
	}
	log.Printf("concat:%+v\n",reply)
	for i:=0;i<10;i++ {
		//将span注册到header中
		req:=injectSpanToCtx(span)
		//设置traceid
		ctx:=metadata.AppendToOutgoingContext(context.Background(),"wfuproject_trace_label",fmt.Sprintf("%d",i),
			"Uber-Trace-Id",req.Header["Uber-Trace-Id"][0])
	    replyS,err:=client.Sum(ctx,&generate.SumRequest{
			A:                    rand.Int63n(1000),
			B:                    rand.Int63n(12022),
		})
		if err!=nil {
			log.Printf("sum error,err:%+v\n",err)
		}
		log.Printf("sum:%+v\n",replyS)
		time.Sleep(time.Millisecond*10)
	}
}

func injectSpanToCtx(span opentracing.Span)*http.Request{
	//将span注册到header中
	req,err:=http.NewRequest("GET","",nil)
	if err!=nil {
		panic(err.Error())
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
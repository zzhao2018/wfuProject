package midware

import (
	"context"
	"google.golang.org/grpc/metadata"
	"wfuProject/logs"
)

/*
 设置traceid的中间件
*/
const(
	TraceLabel="wfuproject_trace_label"
)


func NewPrepareMidWare()MidWare{
	return func(next MidWareFunc) MidWareFunc {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			var traceId string
			//从context中获得traceid
			md,ok:=metadata.FromIncomingContext(ctx)
			if ok==true {
				valueList:=md.Get(TraceLabel)
				if valueList!=nil&&len(valueList)>0 {
					traceId=valueList[0]
				}else{
					traceId=logs.GenTraceId()
				}
			}else{
				traceId=logs.GenTraceId()
			}
			//设置traceid
			ctx=logs.SetTraceIdFromData(ctx,traceId)
			resp,err=next(ctx,req)
			return
		}
	}
}

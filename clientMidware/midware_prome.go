package clientMidware

import (
	"context"
	"time"
	"wfuProject/clientUtil"
	"wfuProject/logs"
	"wfuProject/midware"
)

var promeScanTool =midware.NewServerScanTool("client_request_num","client_error_count","client_method_usetime")

func NewClientPromeMidWare(nextFunc ClientMidwareFunc)ClientMidwareFunc{
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		//统计请求次数
		serverMeta,err:=clientUtil.GetMetaDataFromContext(ctx)
		if err!=nil {
			logs.Error(ctx,"GetMetaDataFromContext error,err:%+v\n",err)
		}
		promeScanTool.IncrRequestTime(serverMeta.ServerName,serverMeta.MethodName)
		//统计耗时
		timeStart:=time.Now()
		resp,err=nextFunc(ctx,request)
		promeScanTool.IncrErrorTime(serverMeta.ServerName,serverMeta.MethodName,err)
		promeScanTool.CalTimeUse(serverMeta.ServerName,serverMeta.MethodName,time.Since(timeStart).Nanoseconds()/1000)
		return
	}
}

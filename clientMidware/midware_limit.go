package clientMidware

import (
	"context"
	"wfuProject/clientUtil"
	"wfuProject/logs"
	"wfuProject/midware"
)

func NewClientLimitMidWare(limit midware.LimitMid)ClientMidware{
	return func(nextFunc ClientMidwareFunc) ClientMidwareFunc {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			//获取元数据
			clientMetaData,errM:=clientUtil.GetMetaDataFromContext(ctx)
			if errM!=nil {
				logs.Error(ctx,"get metadata error,err:%+v\n",err)
				return
			}
			//判断是否能通过
			allow:=limit.Allow()
			if allow==false {
				logs.Error(ctx,"method %s out of limit\n",clientMetaData.MethodName)
				return
			}
			resp,err=nextFunc(ctx,request)
			return
		}
	}
}
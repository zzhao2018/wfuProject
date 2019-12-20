package clientMidware

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"wfuProject/clientUtil"
	"wfuProject/logs"
)

func NewClientFuseMidWare(nextFunc ClientMidwareFunc)ClientMidwareFunc{
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var(
			response interface{}
			err error
		)
		//获得metadata
		serverMetaData,err:=clientUtil.GetMetaDataFromContext(ctx)
		if err!=nil {
			logs.Error(ctx,"midware_fuse NewClientFuseMidWare GetMetaDataFromContext error,err:%+v\n", err)
			return nil, err
		}
		//使用熔断器
		hystrix.ConfigureCommand(serverMetaData.ServerName,hystrix.CommandConfig{
			Timeout:                300,
			MaxConcurrentRequests:  100,
			ErrorPercentThreshold:  25,
		})
		hysErr:=hystrix.Do(serverMetaData.ServerName, func() error {
			response,err=nextFunc(ctx,request)
			return err
		},nil)
		if hysErr!=nil {
			logs.Error(ctx,"midware_fuse NewClientFuseMidWare Do error,err:%+v\n",hysErr)
			return nil,hysErr
		}
		return response,nil
	}
}
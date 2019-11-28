package logs

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type traceIdKey struct {}

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func GetTraceId(ctx context.Context)(string,error){
	v,ok:=ctx.Value(traceIdKey{}).(string)
	if ok==false {
		err:=fmt.Errorf("not found GetTraceId exception")
		return "",err
	}
	return v,nil
}

func SetTraceId(ctx context.Context)context.Context{
	//初始化traceid
	timeN:=time.Now()
	traceid:=fmt.Sprintf("%04d%02d%02d%02d%02d%02d%06d",timeN.Year(),timeN.Month(),timeN.Day(),
		timeN.Hour(),timeN.Minute(),timeN.Second(),rand.Intn(1000000))
	return context.WithValue(ctx,traceIdKey{},traceid)
}

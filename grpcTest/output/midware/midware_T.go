package midware

import (
	"context"
	"fmt"
	"log"
	"time"
	"wfuProject/codeGenerate/output/generate"
)

func MidWareT(){
	mid2:= func(next MidWareFunc)MidWareFunc {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			//强制转换
			_,ok:=req.(*generate.SumRequest)
			if ok==false {
				log.Printf("%+v not sum type\n",req)
				err:=fmt.Errorf("%+v not sum type",req)
				return nil,err
			}
			resp,err=next(ctx,req)
			time.Sleep(time.Second)
			return resp,err
		}
	}


	mid1:= func(next MidWareFunc)MidWareFunc {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			timeN:=time.Now().UnixNano()
			resp,err=next(ctx,req)
			timeAF:=time.Now().UnixNano()
			log.Printf("%+v use time:%+v\n",req,(timeAF-timeN)/1000)
			return
		}
	}
	AddUserMidWare(mid1,mid2)
}
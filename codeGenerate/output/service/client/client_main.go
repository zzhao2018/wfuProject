package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"wfuProject/clientService"
	"wfuProject/codeGenerate/output/generate"
	"wfuProject/codeGenerate/output/service/client/clientTool"
)

func main() {
	for  {
		cd:= clientTool.NewClientDeal("/test/test", clientService.OptClientTraceId("1228334"))
		ctx:=context.Background()
		reply,err:=cd.Concat(ctx,&generate.ConcatRequest{
			Data1:"hello,",
			Data2:"zzhao",
		})
		if err!=nil {
			log.Printf("error:%+v\n",err)
			time.Sleep(time.Millisecond*20)
			fmt.Printf("==============replySub ERROR================\n")
		}else{
			fmt.Printf("data:%+v\n",reply.Result)
			time.Sleep(time.Millisecond*20)
			fmt.Printf("---------------replySub SUCCESS---------------\n\n")
		}
		replySum,err:=cd.Sum(ctx,&generate.SumRequest{
			A:                    1290,
			B:                    125,
		})
		if err!=nil {
			log.Printf("error:%+v\n",err)
			time.Sleep(time.Millisecond*20)
			fmt.Printf("==============replySub ERROR================\n")
		}else{
			fmt.Printf("data:%+v\n",replySum.V)
			time.Sleep(time.Millisecond*20)
			fmt.Printf("---------------replySub SUCCESS---------------\n\n")
		}
		replySub,err:=cd.Sub(ctx,&generate.SumRequest{
			A:                    10,
			B:                    30,
		})
		if err!=nil {
			log.Printf("error:%+v\n",err)
			time.Sleep(time.Millisecond*20)
			fmt.Printf("==============replySub ERROR================\n")
		}else{
			fmt.Printf("data:%+v\n",replySub.V)
			time.Sleep(time.Millisecond*20)
			fmt.Printf("---------------replySub SUCCESS---------------\n\n")
		}
	}
}


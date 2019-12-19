package main

import (
	"context"
	"fmt"
	"log"
	"wfuProject/clientService"
	"wfuProject/codeGenerate/output/generate"
	"wfuProject/codeGenerate/output/service/client/clientTool"
)

func main() {
	for i:=0;i<10 ;i++  {
		cd:= clientTool.NewClientDeal("/test/serverTest", clientService.OptClientTraceId("12282"))
		ctx:=context.Background()
		reply,err:=cd.Concat(ctx,&generate.ConcatRequest{
			Data1:"hello,",
			Data2:"zzhao",
		})
		if err!=nil {
			log.Printf("error:%+v\n",err)
			return
		}
		fmt.Printf("data:%+v\n",reply.Result)
		replySum,err:=cd.Sum(ctx,&generate.SumRequest{
			A:                    1290,
			B:                    125,
		})
		if err!=nil {
			log.Printf("error:%+v\n",err)
			return
		}
		fmt.Printf("data:%+v\n",replySum.V)

		replySub,err:=cd.Sub(ctx,&generate.SumRequest{
			A:                    10,
			B:                    30,
		})
		if err!=nil {
			log.Printf("error:%+v\n",err)
			return
		}
		fmt.Printf("data:%+v\n",replySub.V)
	}

}


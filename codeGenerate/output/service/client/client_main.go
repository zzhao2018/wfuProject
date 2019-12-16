package main

import (
	"context"
	"fmt"
	"log"
	"wfuProject/codeGenerate/output/generate"
	"wfuProject/codeGenerate/output/service/client/clientService"
)

func main() {
	cd:= clientService.NewClientDeal("clientTest1216", clientService.OptClientTraceId("12282"))
	ctx:=context.Background()
	reply,err:=cd.Concat(ctx,&generate.ConcatRequest{
		Data1:"hello,",
		Data2:"wsl",
	})
	if err!=nil {
		log.Printf("error:%+v\n",err)
		return
	}
	fmt.Printf("data:%+v\n",reply.Result)
	replySum,err:=cd.Sum(ctx,&generate.SumRequest{
		A:                    1290,
		B:                    1205,
	})
	if err!=nil {
		log.Printf("error:%+v\n",err)
		return
	}
	fmt.Printf("data:%+v\n",replySum.V)
	
	replySub,err:=cd.Sub(ctx,&generate.SumRequest{
		A:                    1020,
		B:                    30,
	})
	if err!=nil {
		log.Printf("error:%+v\n",err)
		return
	}
	fmt.Printf("data:%+v\n",replySub.V)
}


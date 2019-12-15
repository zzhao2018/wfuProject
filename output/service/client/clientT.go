package main

import (
	"context"
	"fmt"
	"log"
	"wfuProject/output/generate"
	"wfuProject/output/service/client/clinetServer"
)

func main() {
	cd:= clinetServer.NewClientDeal("clientTest", clinetServer.OptClientTraceId("12282"))
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
}


package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"wfuProject/codeGenerate/output/generate"
)

func main() {
	conn,err:=grpc.Dial("localhost:12306",grpc.WithInsecure())
	if err!=nil {
		log.Printf("dial error,err:%+v\n",err)
		return
	}
	client:=generate.NewTestClient(conn)
	reply,err:=client.Concat(context.Background(),&generate.ConcatRequest{
		Data1:                "hello,",
		Data2:                "zzh",
	})
	if err!=nil {
		log.Printf("Concat error,err:%+v\n",err)
	}
	log.Printf("concat:%+v\n",reply)
	replyS,err:=client.Sum(context.Background(),&generate.SumRequest{
		A:                    12003,
		B:                    6120,
	})
	if err!=nil {
		log.Printf("sum error,err:%+v\n",err)
	}
	log.Printf("sum:%+v\n",replyS)

}

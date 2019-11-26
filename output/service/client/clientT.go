package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"time"
	"wfuProject/output/generate"
)

func main() {
	conn,err:=grpc.Dial("localhost:12306",grpc.WithInsecure())
	if err!=nil {
		log.Printf("dial error,err:%+v\n",err)
		return
	}
	client:= generate.NewTestClient(conn)
	reply,err:=client.Concat(context.Background(),&generate.ConcatRequest{
		Data1:                "hello,",
		Data2:                "zzh",
	})
	if err!=nil {
		log.Printf("Concat error,err:%+v\n",err)
	}
	log.Printf("concat:%+v\n",reply)
	for  {
	    _,err:=client.Sum(context.Background(),&generate.SumRequest{
			A:                    rand.Int63n(1000),
			B:                    rand.Int63n(12022),
		})
		if err!=nil {
			log.Printf("sum error,err:%+v\n",err)
		}
	//	log.Printf("sum:%+v\n",replyS)
		time.Sleep(time.Millisecond*10)
	}

}

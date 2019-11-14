package main

import (
	"awesomeProject/videoLearn/part49/model"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//连接rpc
	grpcClient,err:=grpc.Dial("localhost:10244",grpc.WithInsecure())
	if err!=nil {
		log.Printf("dial error,err:%+v\n",err)
		return
	}
	defer grpcClient.Close()
	//初始化model
	client:=model.NewAddClient(grpcClient)
	reply,err:=client.Sum(context.Background(),&model.SumRequest{
		A:                    186,
		B:                    1233,
	})
	if err!=nil {
		log.Printf("get client error,err:%+v\n",err)
		return
	}
	log.Printf("sum result:%+v\n",reply)
	log.Println("===========================")
	replyConcat,err:=client.Concat(context.Background(),&model.ConcatRequest{Data2:"zzhao",Data1:"hello,",})
	if err!=nil {
		log.Printf("get concat error,err:%+v\n",err)
		return
	}
	log.Printf("concat result:%+v\n",replyConcat)
}

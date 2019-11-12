package etcdRegister

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
	"wfuProject/register"
)

// /test/serverTest/127.0.0.1:8808
// /test/serverTest/127.0.0.2:6377
func TestEtcdRegister_Init(t *testing.T) {
	//初始化
	server,err:=register.InitServer(context.TODO(),"etcd",
		register.RegisterInitAddr([]string{"localhost:2379"}),
		register.RegisterInitHeartBeat(5),
		register.RegisterInitRegisterPath("/test"),
		register.RegisterInitTimeOut(time.Second*5))
	if err!=nil {
		log.Printf("main init server error,err:%+v\n",err)
		return
	}
	//
	serverData:=&register.Server{
		Name: "serverTest",
		Node: make([]*register.ServerNode,0),
	}
	serverData.Node=append(serverData.Node,&register.ServerNode{Ip:"127.0.0.1",Port:"8808"})
	serverData.Node=append(serverData.Node,&register.ServerNode{Ip:"127.0.0.2",Port:"6377"})
	server.Register(context.TODO(),serverData)

	//拉取数据
	go func() {
		time.Sleep(time.Second*8)
		fmt.Println("=============begin add new node=============")
		serverData:=&register.Server{
			Name: "serverTest",
			Node: make([]*register.ServerNode,0),
		}
		serverData.Node=append(serverData.Node,&register.ServerNode{Ip:"192.23.121.13",Port:"8800"})
		serverData.Node=append(serverData.Node,&register.ServerNode{Ip:"192.66.127.22",Port:"2667"})
		server.Register(context.TODO(),serverData)
	}()
	for  {
		serverResp,err:=server.GetServer(context.TODO(),"serverTest")
		if err!=nil {
			t.Fatalf("get data error,err:%+v\n",err)
			return
		}
		for _,ele:=range serverResp.Node {
			fmt.Printf("get node:%+v\n",ele)
		}
		fmt.Printf("=====================\n")
		time.Sleep(time.Second*2)
	}
}


func TestEtcdRegister_CheckServer(t *testing.T) {
	loc:=rand.Intn(1)
	t.Logf("data:%+v\n",loc)
}
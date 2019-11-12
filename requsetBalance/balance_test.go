package requsetBalance

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
	"wfuProject/register"
)

var(
	coundown sync.WaitGroup
	key="datakey"
	value="dataValue"
)

func goWork(ctx context.Context){
	time.Sleep(time.Second*2)
	fmt.Println("work begin....")
	data:=ctx.Value(key).(string)
	fmt.Printf("gotine get(%s) data:%s\n",key,data)
	coundown.Done()
}

func TestNewPollBalance(t *testing.T) {

	ctx:=context.WithValue(context.Background(),key,value)
	coundown.Add(1)
	go goWork(ctx)
	coundown.Wait()

}

func TestRandomBalance_Select(t *testing.T) {
	//产生结点
	nodeList:=[]*register.ServerNode{}
	weight:=[]int{200,300,300,200}
	for i:=0;i<4 ;i++  {
		nodeList=append(nodeList,&register.ServerNode{
			Ip:     fmt.Sprintf("127.0.0.%d",i),
			Port:   "6379",
			Weight: weight[i%len(weight)],
		})
	}
	//发送请求
	balance,err:=GetBalance(B_PollingWeightBalance)
	if err!=nil {
		t.Fatalf("get balance error,err:%+v\n",err)
		return
	}
	countMap:=make(map[string]int,0)
	for i:=0;i<1000 ;i++  {
		node:=balance.Select(context.TODO(),nodeList)
		name:=fmt.Sprintf("%s:%s",node.Port,node.Ip)
		countMap[name]++
	}
	for name,value:=range countMap {
		fmt.Printf("ip:%s num:%d\n",name,value)
	}
}

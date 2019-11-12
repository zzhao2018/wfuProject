package requsetBalance

import (
	"context"
	"math/rand"
	"wfuProject/register"
)

//随机
type RandomBalance struct {
}

func NewRandomBalance()*RandomBalance{
	return &RandomBalance{}
}

func(r *RandomBalance)Select(ctx context.Context,nodeList []*register.ServerNode)(*register.ServerNode){
	if len(nodeList)<=0 {
		return nil
	}
	//产生随机数
	loc:=rand.Intn(len(nodeList))
	//获取下标
	return nodeList[loc]
}

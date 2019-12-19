package requsetBalance

import (
	"context"
	"math/rand"
	"time"
	"wfuProject/register"
)

const(
	R_DefaulfWeight=100
)

//加权随机
type RandomWeightBalance struct {
}

func NewRandomWeightBalance()*RandomWeightBalance{
	return &RandomWeightBalance{}
}

func(r *RandomWeightBalance)Select(ctx context.Context,nodeList []*register.ServerNode)*register.ServerNode{
	if len(nodeList)<=0 {
		return nil
	}
	//整合权值
	allWeight:=0
	for _,ele:=range nodeList {
		if ele.Weight<=0 {
			ele.Weight=R_DefaulfWeight
		}
		allWeight+=ele.Weight
	}
	//产生随机数
	rand.Seed(time.Now().UnixNano())
	randWeight:=rand.Intn(allWeight)
	//判断权值所属结点
	for _,ele:=range nodeList {
		randWeight-=ele.Weight
		if randWeight<0 {
			return ele
		}
	}
	return nil
}

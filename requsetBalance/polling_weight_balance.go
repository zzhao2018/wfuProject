package requsetBalance

import (
	"context"
	"wfuProject/register"
)

type PollingWeightBalance struct {
	index int
}

func NewPollingWeightBalance()Balance{
	return &PollingWeightBalance{
		index:0,
	}
}

func(p *PollingWeightBalance)Select(ctx context.Context,nodeList []*register.ServerNode)*register.ServerNode{
	if(len(nodeList)<=0){
		return nil
	}
	//计算权重之和
	allIndex:=0
	for _,ele:=range nodeList {
		if ele.Weight<=0 {
			ele.Weight=R_DefaulfWeight
		}
		allIndex+=ele.Weight
	}
	indexCopy:=p.index
	if p.index>=allIndex {
		p.index=0
		return nodeList[p.index]
	}
	//确定对应元素
	for _,ele:=range nodeList {
		indexCopy-=ele.Weight
		if indexCopy<0 {
			//更新
			p.index=(p.index+1)%allIndex
			return ele
		}
	}
	return nil
}

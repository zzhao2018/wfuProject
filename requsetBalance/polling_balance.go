package requsetBalance

import (
	"context"
	"wfuProject/register"
)

type PollBalance struct {
	balanceId int
}

func NewPollBalance()Balance{
	return &PollBalance{
		balanceId:0,
	}
}


func(p *PollBalance)Select(ctx context.Context,nodeList []*register.ServerNode)*register.ServerNode{
	if len(nodeList)<=0 {
		return nil
	} else if(p.balanceId>=len(nodeList)){
		p.balanceId=0
		return nodeList[0]
	}
	//获取数据
	selectData:=nodeList[p.balanceId]
	//更新下标
	p.balanceId=(p.balanceId+1)%len(nodeList)
	return selectData
}

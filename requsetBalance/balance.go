package requsetBalance

import (
	"context"
	"fmt"
	"log"
	"wfuProject/register"
)

type Balance interface {
	Select(ctx context.Context,nodeList []*register.ServerNode)*register.ServerNode
}

const(
	B_PollingBalance = iota
	B_PollingWeightBalance
	B_RandomBalance
	B_RandomWeightBalance
)


func GetBalance(balancetype int)(Balance,error){
	switch balancetype {
	case B_PollingBalance:
		return NewPollBalance(),nil
	case B_PollingWeightBalance:
		return NewPollingWeightBalance(),nil
	case B_RandomBalance:
		return NewRandomBalance(),nil
	case B_RandomWeightBalance:
		return NewRandomWeightBalance(),nil
	default:
		log.Printf("balance type %d not found\n",balancetype)
		err:=fmt.Errorf("balance type not found")
		return nil,err
	}
}
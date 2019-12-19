package clientMidware

import (
	"context"
	"fmt"
	"log"
	"wfuProject/clientUtil"
	"wfuProject/register"
	"wfuProject/requsetBalance"
)

func NewClientBalanceMidWare(balance requsetBalance.Balance)ClientMidware{
	return func(nextFunc ClientMidwareFunc) ClientMidwareFunc {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			//获取元数据
			var serverMetaData *clientUtil.ClientMetaData
			serverMetaData,err=clientUtil.GetMetaDataFromContext(ctx)
			if err!=nil {
				log.Printf("NewClientBalanceMidWare GetMetaDataFromContext error,err:%+v\n",err)
				return
			}
			//负载均衡
			var canUseList []*register.ServerNode
			usedMap:=make(map[*register.ServerNode]bool,len(serverMetaData.NodeList))
			canUseList,err=getCanUseList(usedMap,serverMetaData.NodeList)
			if err!=nil {
				log.Printf("midware_balance NewClientBalanceMidWare getCanUseList error,err:%+v\n",err)
				return
			}
			for{
				//无结点可用，则退出
				if len(canUseList)==0 {
					err=fmt.Errorf("midware_balance no node can use error")
					log.Printf("midware_balance no node can use error\n")
					return
				}
				//利用负载均衡算法获取结点
				selectNode:=balance.Select(ctx,canUseList)
				//更新
				usedMap[selectNode]=true
				canUseList,err=getCanUseList(usedMap,canUseList)
				if err!=nil {
					log.Printf("midware_balance NewClientBalanceMidWare getCanUseList error,err:%+v\n",err)
					return
				}
				//传递
				serverMetaData.SelectNode=selectNode
				response,err=nextFunc(ctx,request)
				if err!=nil {
					log.Printf( "midware_balance NewClientBalanceMidWare request node (%+v) error,err:%+v\n",selectNode,err)
					continue
				}else {
					return
				}
			}
		}
	}
}

func getCanUseList(nodeMap map[*register.ServerNode]bool,nodeList []*register.ServerNode)([]*register.ServerNode,error){
	if len(nodeList)<=0 {
		err:=fmt.Errorf("midware_balance getCanUseList error")
		return nil,err
	}
	var resultList=make([]*register.ServerNode,0)
	for _,ele:=range nodeList {
		if _,ok:=nodeMap[ele];ok==false {
			resultList=append(resultList,ele)
		}
	}
	return resultList,nil
}
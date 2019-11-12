package register

import (
	"context"
	"fmt"
	"log"
	"sync"
)

/*
  用于管理注册引擎的插件
 使用map保存register插件
*/
var(
	registerMgr *RegisterMgr=&RegisterMgr{
		registerMap: make(map[string]Register,8),
		lock:        sync.Mutex{},
	}
)

type RegisterMgr struct{
	registerMap map[string]Register
	lock sync.Mutex
}

//将服务注册引擎注册入mgr
func(r *RegisterMgr)registerServer(register Register)error{
	r.lock.Lock()
	defer r.lock.Unlock()
	//判断register是否存在
	if _,ok:=r.registerMap[register.Name()];ok==true{
		log.Printf("registerServer error,register double\n")
		err:=fmt.Errorf("register double")
		return err
	}
	r.registerMap[register.Name()]=register
	return nil
}

//服务注册引擎初始化
func(r *RegisterMgr)initServer(ctx context.Context,serverName string,optFunc ...RegisterOptionFunc)(Register,error){
	r.lock.Lock()
	defer r.lock.Unlock()
	//从map中获得服务
	v,ok:=r.registerMap[serverName]
	if ok==false {
		log.Printf("initServer error,no server in map\n")
		err:=fmt.Errorf("no server in map")
		return nil,err
	}
	//初始化服务
	err:=v.Init(ctx,optFunc...)
	if err!=nil {
		log.Printf("initServer error,err:%+v\n",err)
		return nil,err
	}
	return v,nil
}

func RegisterServer(register Register)error  {
	return registerMgr.registerServer(register)
}

func InitServer(ctx context.Context,serverName string,optFunc ...RegisterOptionFunc)(Register,error){
	return registerMgr.initServer(ctx,serverName,optFunc...)
}
package main

import (
	"fmt"
	"log"
)

/*
  生成器管理中心
*/
const DirGenName="dirGenerator"

var GenMgrInstance *GenMgr=&GenMgr{
	genMap:make(map[string]Generator),
}

type GenMgr struct {
	genMap map[string]Generator
}

//将生成器注册到mgr里
func(g *GenMgr)RegisterGen(genName string,gen Generator)error{
	if _,ok:=g.genMap[genName];ok==true {
		err:=fmt.Errorf("generator exist exception")
		return err
	}
	g.genMap[genName]=gen
	return nil
}


func(g *GenMgr)Run(opt *GenOption)error{
	//遍历所有生成器，调用生成器的run方法
	var err error
	//产生dir
	dirGen,ok:=g.genMap[DirGenName]
	if ok==false {
		log.Printf("Run get dir generator error\n")
		err=fmt.Errorf("get dir generator error")
		return err
	}
	err=dirGen.Run(opt)
	if err!=nil {
		log.Printf("Run generator dir run error,err:%+v\n",err)
		return err
	}
	//产生其余生成器
	for name,eleValue:=range g.genMap {
		if name==DirGenName {
			continue
		}
		err=eleValue.Run(opt)
		if err!=nil {
			log.Printf("Run generator %s run error,err:%+v\n",name,err)
			return err
		}
	}
	return nil
}

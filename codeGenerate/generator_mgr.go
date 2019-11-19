package main

import (
	"fmt"
	"github.com/emicklei/proto"
	"log"
	"os"
)

/*
  生成器管理中心
*/
const DirGenName="dirGenerator"

var GenMgrInstance *GenMgr=&GenMgr{
	genMap:make(map[string]Generator),
	protoData:&protoMetaData{
		Rpc:make([]*proto.RPC,0),
		Service:&proto.Service{},
		Messages:make([]*proto.Message,0),
	},
}

type GenMgr struct {
	genMap map[string]Generator
	protoData *protoMetaData
}

//保存protobuf元数据
type protoMetaData struct {
	Rpc []*proto.RPC
	Service *proto.Service
	Messages []*proto.Message
}

/*****************抽取proto元数据****************/
func(g *GenMgr)parseProto(opt *GenOption)error{
	//打开文件
	fileData,err:=os.Open(opt.ProtoFilePath)
	if err!=nil {
		log.Printf("generator controller run open file error,err:%+v\n",err)
		return err
	}
	defer fileData.Close()
	//解析protobuf
	praseData:=proto.NewParser(fileData)
	protoData,err:=praseData.Parse()
	if err!=nil {
		log.Printf("generator_mgr parseProto prase protobuf error,err:%+v\n",err)
		return err
	}
	//获取protobuf
	proto.Walk(protoData,proto.WithService(g.dealService),proto.WithRPC(g.dealRpc),proto.WithMessage(g.dealMessage))
	return nil
}

func(g *GenMgr)dealService(service *proto.Service){
	g.protoData.Service=service
}

func(g *GenMgr)dealRpc(rpc *proto.RPC){
	g.protoData.Rpc=append(g.protoData.Rpc,rpc)
}

func(g *GenMgr)dealMessage(message *proto.Message){
	g.protoData.Messages=append(g.protoData.Messages,message)
}


/*************注册、运行****************/
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
	var err error
	//抽取元数据
	err=g.parseProto(opt)
	if err!=nil {
		log.Printf("generator_mgr Run error,err:%+v\n",err)
		return err
	}
	//产生dir
	dirGen,ok:=g.genMap[DirGenName]
	if ok==false {
		log.Printf("Run get dir generator error\n")
		err=fmt.Errorf("get dir generator error")
		return err
	}
	err=dirGen.Run(opt,g.protoData)
	if err!=nil {
		log.Printf("Run generator dir run error,err:%+v\n",err)
		return err
	}
	//遍历所有生成器，调用生成器的run方法
	for name,eleValue:=range g.genMap {
		if name==DirGenName {
			continue
		}
		err=eleValue.Run(opt,g.protoData)
		if err!=nil {
			log.Printf("Run generator %s run error,err:%+v\n",name,err)
			return err
		}
	}
	return nil
}

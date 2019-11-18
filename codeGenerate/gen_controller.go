package main

import (
	"github.com/emicklei/proto"
	"io/ioutil"
	"log"
	"os"
	"fmt"
	"path"
	"strings"
)

const(
	serverTemplePath=`./temple/grpcServer.temple`
	generatorPath=`service/server`
)

/*
  自动生成rpc代码
*/
type GenController struct {
	rpc []*proto.RPC
	service *proto.Service
	messages []*proto.Message
}

func init(){
	genController:=&GenController{}
	err:=GenMgrInstance.RegisterGen("controllerGen",genController)
	if err!=nil {
		log.Printf("register generator error,err:%+v\n",err)
		return
	}
}

func(g *GenController)Run(opt *GenOption)error{
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
		log.Printf("controller Run prase protobuf error,err:%+v\n",err)
		return err
	}
	//获取protobuf
	proto.Walk(protoData,proto.WithService(g.dealService),proto.WithRPC(g.dealRpc),proto.WithMessage(g.dealMessage))
	//构建服务端代码
	err=g.buildServer(opt)
	if err!=nil{
		log.Printf("controller Run buildServer error,err:%+v\n",err)
		return err
	}
	return nil
}

//生成服务端代码
func(g *GenController)buildServer(opt *GenOption)error{
	//打开需生成的代码文件
	fileName:=path.Join(opt.OutputPath,generatorPath,fmt.Sprintf("%s.go",g.service.Name))
	file,err:=os.OpenFile(fileName,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0755)
	if err!=nil {
		log.Printf("controller Run open file error,err:%+v\n",err)
		return err
	}
	defer file.Close()
	//打开模板
	templeDataByte,err:=ioutil.ReadFile(serverTemplePath)
	//替换server名称
	serverReplaceData:=strings.ReplaceAll(string(templeDataByte),"{{ServerName}}",g.service.Name)
	fmt.Fprintf(file,serverReplaceData)
	fmt.Fprintf(file,"\n\n")
	//写入方法
	for _,method:=range g.rpc {
		fmt.Fprintf(file,
			"func(s *%s)%s(ctx context.Context, req *generate.%s) (*generate.%s, error){\n\treturn nil,nil\n}\n\n",
			g.service.Name,method.Name,method.RequestType,method.ReturnsType)
	}
	return nil
}



/**************获取proto解析信息***************/
func(g *GenController)dealService(service *proto.Service){
	g.service=service
}

func(g *GenController)dealRpc(rpc *proto.RPC){
	g.rpc=append(g.rpc,rpc)
}

func(g *GenController)dealMessage(message *proto.Message){
	g.messages=append(g.messages,message)
}


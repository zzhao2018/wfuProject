package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

const(
	generatorPath=`service/server`
)

/*
  自动生成rpc代码
*/
type GenController struct {
}

func init(){
	genController:=&GenController{}
	err:=GenMgrInstance.RegisterGen("controllerGen",genController)
	if err!=nil {
		log.Printf("register generator error,err:%+v\n",err)
		return
	}
}

func(g *GenController)Run(opt *GenOption,metaData *protoMetaData)error{
	//构建服务端代码
	err:=g.buildServer(opt,metaData)
	if err!=nil{
		log.Printf("controller Run buildServer error,err:%+v\n",err)
		return err
	}
	return nil
}

//生成服务端代码
func(g *GenController)buildServer(opt *GenOption,metaData *protoMetaData)error{
	midMap:=make(map[string]interface{})
	midMap["OutputPath"]=metaData.OutputPath
	for _,rpcEle:=range metaData.Rpc {
		midMap["Rpc"]=rpcEle
		//打开需生成的代码文件
		fileName:=path.Join(opt.OutputPath,generatorPath,fmt.Sprintf("%sController.go",rpcEle.Name))
		file,err:=os.OpenFile(fileName,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0755)
		if err!=nil {
			log.Printf("controller Run open file error,err:%+v\n",err)
			return err
		}
		defer file.Close()
		//解析template
		err=parseTemple(file,templeData,midMap)
		if err!=nil {
			log.Printf("gen controller parseTemple error,err:%+v\n",err)
			return err
		}
	}
	return nil
}
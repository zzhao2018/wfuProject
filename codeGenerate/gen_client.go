package main

import (
	"log"
	"os"
	"path"
	"wfuProject/util"
)

type GenClient struct {
}

const(
	genclientName="genClient"
	genClientServiceDir="service/client/clientTool"
	genClientDir="service/client"
)

func init(){
	genClient:=&GenClient{}
	err:=GenMgrInstance.RegisterGen(genclientName,genClient)
	if err!=nil {
		log.Printf("init %s RegisterGen error,err:%+v\n",err)
		return
	}
}

func(g *GenClient) Run(opt *GenOption,metaData *protoMetaData)error{
	/*********创建client的main函数*******/
	//构造main
	paths:=path.Join(opt.OutputPath,genClientDir,"client_main.go")
	//存在文件,跳过
	if util.CheckFileExist(paths)==false{
		err:=createFile(paths,templeClientMain,nil)
		if err!=nil {
			log.Printf("gen_client Run create main file error,err:%+v\n",err)
			return err
		}
	}
	//构造mainserver
	pathServiceS:=path.Join(opt.OutputPath,genClientServiceDir,"client_tool.go")
	err:=createFile(pathServiceS,templeClientService,metaData)
	if err!=nil {
		log.Printf("gen_client Run create service file error,err:%+v\n",err)
		return err
	}
	return nil
}

//产生函数
func createFile(paths string,templeStr string,metaData *protoMetaData)error{
	fileF,err:=os.OpenFile(paths,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0755)
	if err!=nil {
		log.Printf("gen_client Run openfile error,err:%+v\n",err)
		return err
	}
	defer fileF.Close()
	//解析文本
	err=parseTemple(fileF,templeStr,metaData)
	if err!=nil {
		log.Printf("gen_client Run parseTemple error,err:%+v\n",err)
		return err
	}
	return nil
}
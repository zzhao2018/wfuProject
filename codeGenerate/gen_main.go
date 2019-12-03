package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"wfuProject/util"
)

type GenMain struct {
}

const(
	genMainName="genMain"
	mainGeneratorPath=`service/server`
)

func init(){
	genMain:=&GenMain{}
	err:=GenMgrInstance.RegisterGen(genMainName,genMain)
	if err!=nil {
		log.Printf("genMain registergen error,err:%+v\n",err)
		return
	}
}

func(g *GenMain)Run(opt *GenOption,metaData *protoMetaData)error{
	//构建地址
	pathS:=path.Join(opt.OutputPath,mainGeneratorPath,fmt.Sprintf("%s_main.go",metaData.Service.Name))
	//判断文件是否存在
	if util.CheckFileExist(pathS)==true {
		return nil
	}
	fileF,err:=os.OpenFile(pathS,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0755)
	if err!=nil {
		log.Printf("gen main open file error,err:%+v\n",err)
		return err
	}
	defer fileF.Close()
	//解析文件
	err=parseTemple(fileF,templeMain,metaData)
	if err!=nil {
		log.Printf("gen main parseTemple error,err:%+v\n",err)
		return err
	}
	return nil
}

package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

type GenRouter struct {
}

const(
	genName="routerGen"
	genFileDir="router"
)

//注册模块
func init(){
	gen:=&GenRouter{}
	err:=GenMgrInstance.RegisterGen(genName,gen)
	if err!=nil {
		log.Printf("GenRouter init %s error,err:%+v\n",genName,err)
		return
	}
}


//构建
func(g *GenRouter)Run(opt *GenOption,metaData *protoMetaData)error{
	//创建文件
	genFilePath:=path.Join(opt.OutputPath,genFileDir,fmt.Sprintf("%s_router.go",metaData.Service.Name))
	fileF,err:=os.OpenFile(genFilePath,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0755)
	if err!=nil {
		log.Printf("gen_router Run open file error,err:%+v\n",err)
		return err
	}
	defer fileF.Close()
	//解析template
	err=parseTemple(fileF,routerTemp,metaData)
	if err!=nil {
		log.Printf("gen_router Run parseTemple error,err:%+v\n",err)
		return err
	}
	return nil
}
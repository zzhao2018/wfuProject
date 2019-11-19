package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
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
	//打开需生成的代码文件
	fileName:=path.Join(opt.OutputPath,generatorPath,fmt.Sprintf("%s.go",metaData.Service.Name))
	file,err:=os.OpenFile(fileName,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0755)
	if err!=nil {
		log.Printf("controller Run open file error,err:%+v\n",err)
		return err
	}
	defer file.Close()
	//替换server名称
	serverReplaceData:=strings.ReplaceAll(templeData,"{{ServerName}}",metaData.Service.Name)
	fmt.Fprintf(file,serverReplaceData)
	fmt.Fprintf(file,"\n\n")
	//写入方法
	for _,method:=range metaData.Rpc {
		fmt.Fprintf(file,
			"func(s *%s)%s(ctx context.Context, req *generate.%s) (*generate.%s, error){\n\treturn nil,nil\n}\n\n",
			metaData.Service.Name,method.Name,method.RequestType,method.ReturnsType)
	}
	return nil
}
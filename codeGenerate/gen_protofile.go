package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

type GenProto struct {

}

func init(){
	protoGen:=&GenProto{}
	GenMgrInstance.RegisterGen(protoGen.Name(),protoGen)
}


func(p *GenProto)Name()string{
	return "protoGenerator"
}



//调用系统调用，生成proto代码
func(p *GenProto) Run(opt *GenOption)error{
	//初始化路径
	outputPath:=path.Join(opt.OutputPath,"generate")
	outputArgs:=fmt.Sprintf("--go_out=plugins=grpc:%s",outputPath)
	//执行
	cmd:=exec.Command("protoc",outputArgs,opt.ProtoFilePath)
	cmd.Stderr=os.Stderr
	cmd.Stdout=os.Stdout
	err:=cmd.Run()
	if err!=nil {
		log.Printf("proto cmd Run error,err:%+v\n",err)
		return err
	}
	return nil
}
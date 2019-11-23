package main

import (
	"log"
	"os"
	"path"
)

type GenMidWare struct {
}

const(
	midwareGenName="genMidware"
	midwareSaveDir=`midware`
)

func init(){
	g:=&GenMidWare{}
	err:=GenMgrInstance.RegisterGen(midwareGenName,g)
	if err!=nil {
		log.Printf("%s register generator error,err:%+v\n",midwareGenName,err)
		return
	}
}

func(g *GenMidWare)Run(opt *GenOption,metaData *protoMetaData)error{
	//打开文件
	filename:=path.Join(opt.OutputPath,midwareSaveDir,"midware_center.go")
	fileF,err:=os.OpenFile(filename,os.O_TRUNC|os.O_CREATE|os.O_WRONLY,0755)
	if err!=nil {
		log.Printf("GenMidWare Run error,err:%+v\n",err)
		return err
	}
	defer fileF.Close()
	err=parseTemple(fileF,midwareTemp,nil)
	if err!=nil {
		log.Printf("GenMidWare Run parse temple error,err:%+v\n",err)
		return err
	}
	return nil
}

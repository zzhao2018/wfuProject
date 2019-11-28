package logs

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type LogFile struct {
	file *os.File
	fileDir string
	fileName string
	timeCur string
}

//初始化文本日志
func NewLogFile(fileDir string,fileName string)(*LogFile,error){
	timeN:=time.Now()
	timeCurS:=buildTimeStr(timeN)
	fileF,err:=initFile(fileDir,fileName,timeCurS)
	if err!=nil {
		log.Printf("log NewLogFile initFile error,err:%+v\n",err)
		return nil,err
	}
	return &LogFile{
		fileDir:fileDir,
		fileName:fileName,
		file:fileF,
		timeCur:timeCurS,
	},nil
}

//关闭
func(l *LogFile)Close(){
	l.file.Close()
}

//写入
func(l *LogFile)Write(mess *LogMessMeta)error{
	//测试是否需要切割日志
	err:=l.splitFile()
	if err!=nil {
		log.Printf("log splitFile error,err:%+v\n",err)
		return err
	}
	dataByte,err:=mess.ToByte()
	if err!=nil {
		log.Printf("log LogFile Write ToByte error,err:%+v\n",err)
		return err
	}
	_,err=l.file.Write(dataByte)
	if err!=nil {
		log.Printf("log LogFile Write File error,err:%+v\n",err)
		return err
	}
	return nil
}

//构造日期
func buildTimeStr(timeN time.Time) string {
	return fmt.Sprintf("%04d%02d%02d%02d",timeN.Year(),timeN.Month(),timeN.Day(),timeN.Hour())
}

//初始化文件
func initFile(fileDir string,fileName string,timeCurS string)(*os.File,error){
	filePath:=fmt.Sprintf("%s_%s",timeCurS,fileName)
	filePath=path.Join(fileDir,filePath)
	fileF,err:=os.OpenFile(filePath,os.O_WRONLY|os.O_CREATE|os.O_APPEND,0755)
	if err!=nil {
		log.Printf("NewLogFile error,err:%+v\n",err)
		return nil,err
	}
	return fileF,nil
}

func(l *LogFile)splitFile()error{
	timeN:=time.Now()
	timeNStr:=buildTimeStr(timeN)
	if l.timeCur==timeNStr {
		return nil
	}
	//超时，更新日志
	var err error
	l.file.Close()
	l.file,err=initFile(l.fileDir,l.fileName,timeNStr)
	if err!=nil {
		log.Printf("log splitFile initFile error,err:%+v\n",err)
		return err
	}
	l.timeCur=timeNStr
	return nil
}
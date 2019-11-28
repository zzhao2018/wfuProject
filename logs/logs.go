package logs

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type LogMgr struct {
	messChan chan *LogMessMeta      //用以保存日志数据
	output []LogOutput
	serverName string
	outputLevel LogLevel
	countDownLanch sync.WaitGroup
}

var(
	logMgr *LogMgr
	once sync.Once
	defaultOutput LogOutput=NewLogConsole()
	defaultServerName="defaultServer"
)

//初始化
func InitLog(chansize int,logOutputLevel LogLevel,serverName string){
	if chansize<=0 {
		chansize=defaultMessChanSize
	}
	//初始化日志配置
	once.Do(func() {
		logMgr=&LogMgr{
			messChan:make(chan *LogMessMeta,chansize),
			output:make([]LogOutput,0),
			outputLevel:logOutputLevel,
			serverName:serverName,
		}
		//开启写协程
		go logMgr.run()
		//用于等待优雅关闭
		logMgr.countDownLanch.Add(1)
	})
}

/*******添加输出接口***********/
func LogAddOutPut(outputEle LogOutput){
	if logMgr==nil {
		InitLog(defaultMessChanSize,L_Debug,defaultServerName)
	}
	logMgr.output=append(logMgr.output,outputEle)
}

/************优雅关闭**********/
func LogStop(){
	if logMgr==nil {
		log.Printf("log Stop error,err:logMgr is nil,can not stop\n")
		return
	}
	close(logMgr.messChan)
	logMgr.countDownLanch.Wait()
	for _,outputEle:=range logMgr.output {
		outputEle.Close()
	}
	//更新
	once=sync.Once{}
	logMgr=nil
}

/**************后台写日志方法********************/
func(l *LogMgr)run(){
	for dataMeta:=range l.messChan {
		//若未初始化output，则使用控制台输出
		if l.output==nil||len(l.output)==0 {
			defaultOutput.Write(dataMeta)
			continue
		}
		//写入
		for _,outputEle:=range l.output {
			outputEle.Write(dataMeta)
		}
	}
	//优雅关闭时，等待所有日志写入才关闭
	l.countDownLanch.Done()
}



/**************日志记录方法****************/
func Debug(ctx context.Context,format string,opts... interface{}){
	writeLog(ctx,L_Debug,format,opts...)
}

func Info(ctx context.Context,format string,opts... interface{}){
	writeLog(ctx,L_Info,format,opts...)
}

func Warn(ctx context.Context,format string,opts... interface{}){
	writeLog(ctx,L_Warn,format,opts...)
}

func Error(ctx context.Context,format string,opts... interface{}){
	writeLog(ctx,L_Error,format,opts...)
}

func writeLog(ctx context.Context,logLevel LogLevel,format string,opts... interface{}){
	//判断是否低于level
	if logLevel<logMgr.outputLevel {
		return
	}
	//初始化日志元元素
	traceId,err:=GetTraceId(ctx)
	if err!=nil {
		log.Printf("writeLog GetTraceId error,err:%+v\n",err)
		return
	}
	dataMeta:=&LogMessMeta{
		mess:fmt.Sprintf(format,opts...),
		level:logLevel.GetLogLevelName(),
		timeStr:time.Now().Format("2006-01-02 15:04:05.99"),
		serverName:logMgr.serverName,
		traceId:traceId,
		messLoc:GetMessLoc(),
	}
	//元素放入chan
	select {
	case logMgr.messChan<-dataMeta:
	default:
		return
	}
}





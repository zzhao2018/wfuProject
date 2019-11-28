package logs

import (
	"context"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestInitLog(t *testing.T) {
	//初始化日志
	InitLog(100,L_Debug,"test")
	//初始化文件
	fileDir:=`C:\Users\35278\Desktop\测试数据\go_log_test`
	fileName:=`test.log`
	logF,err:=NewLogFile(fileDir,fileName)
	if err!=nil {
		log.Printf("error,err:%+v\n",err)
		return
	}
	LogAddOutPut(logF)
	ctx:=context.Background()
	ctx=SetTraceId(ctx)
	rand.Seed(time.Now().UnixNano())
	for  {
		i:=rand.Intn(100000)
		Debug(ctx,"error,err:%+v\n",i)
		time.Sleep(time.Millisecond*10)
	}
	LogStop()
	//time.Sleep(time.Second*2)
}

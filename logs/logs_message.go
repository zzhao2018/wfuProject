package logs

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

type LogMessMeta struct {
	mess string
	level string
	timeStr string
	serverName string
	traceId string
	messLoc string
}

func(l *LogMessMeta)ToByte()([]byte,error){
	var buff bytes.Buffer
	_,err:=buff.WriteString(fmt.Sprintf("%s [%s] (%s;%s;%s) %s",l.timeStr,l.level,l.traceId,l.serverName,l.messLoc,l.mess))
	if err!=nil {
		log.Printf("LogMessMeta ToByte error,err:%+v\n",err)
		return nil,err
	}
	return buff.Bytes(),nil
}

func GetMessLoc()string{
	_,file,line,_:=runtime.Caller(3)

	return fmt.Sprintf("%s:%d",filepath.Base(file),line)
}
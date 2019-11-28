package logs

import (
	"log"
	"os"
)

type LogConsole struct {
}

func NewLogConsole()*LogConsole{
	return &LogConsole{}
}

func(l *LogConsole)Close(){
}

func(l *LogConsole)Write(mess *LogMessMeta)error{
	dataByte,err:=mess.ToByte()
	if err!=nil {
		log.Printf("log LogConsole Write ToByte error,err:%+v\n",err)
		return err
	}
	os.Stdout.Write(dataByte)
	return nil
}

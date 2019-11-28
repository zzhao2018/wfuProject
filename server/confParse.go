
package server

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Port  int  `yaml:"port"`
	Prometheus PromeConf `yaml:"prometheus"`
	Limit LimitConf `yaml:"limit"`
	Logs LogsConf ``
}

type LogsConf struct {
	ChanSize int `yaml:"chansize"`
	LogLevel string `yaml:"loglevel"`
	ServerName string `yaml:"servername"`
}

type LimitConf struct {
	Switch_on bool `yaml:"switch_on"`
	Qbs float64 `yaml:"qps"`
	Size float64 `yaml:"size"`
	Type string `yaml:"type"`
	TimeDiff int64 `yaml:"timediff"`
}

type PromeConf struct {
	Switch_on bool `yaml:"switch_on"`
	Port int `yaml:"port"`
}

var conf Conf

const(
	G_TestConfName="test"
	G_ProductConfName="product"
	G_TestConfPath="../../conf/test/test.yaml"
	G_ProductConfPath="../../conf/product/product.yaml"
)


func ParseConfInit(confType string)error{
	//打开文件
	var(
		fileF []byte
		err error
	)
	if confType== G_TestConfName{
		fileF,err=ioutil.ReadFile(G_TestConfPath)
	}else if confType==G_ProductConfName{
		fileF,err=ioutil.ReadFile(G_ProductConfPath)
	}else{
        err=fmt.Errorf("file type not found exception")
	}
	if err!=nil {
		log.Printf("ParseConfInit read file error,err:%+v\n",err)
		return err
	}
	err=yaml.Unmarshal(fileF,&conf)
	if err!=nil {
		log.Printf("ParseConfInit umarshal error,err:%+v\n", err)
		return err
	}
    return nil
}

func GetParseConfPort()int{
	return conf.Port
}

func GetParseConfPrometheus()PromeConf{
	return conf.Prometheus
}

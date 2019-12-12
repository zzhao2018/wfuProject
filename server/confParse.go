package server

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Conf struct {
	ServerName string       `yaml:"servername"`
	Port       int          `yaml:"port"`
	Prometheus PromeConf    `yaml:"prometheus"`
	Limit      LimitConf    `yaml:"limit"`
	Logs       LogsConf     `yaml:"logs"`
	Register   RegisterConf `yaml:"register"`
	Trace      TraceConf    `yaml:"trace"`
}

type TraceConf struct {
	Switch_on   bool    `yaml:"switch_on"`
	Report_addr string  `yaml:"report_addr"`
	Sample_type string  `yaml:"sample_type"`
	Sample_rate float64 `yaml:"sample_rate"`
}

type RegisterConf struct {
	Switch_on    bool          `yaml:"switch_on"`
	Addr         []string      `yaml:"addr"`
	TimeOut      time.Duration `yaml:"timeOut"`
	RegisterPath string        `yaml:"registerPath"`
	HeartBeat    int64         `yaml:"heartBeat"`
}

type LogsConf struct {
	ChanSize int    `yaml:"chansize"`
	LogLevel string `yaml:"loglevel"`
}

type LimitConf struct {
	Switch_on bool    `yaml:"switch_on"`
	Qbs       float64 `yaml:"qps"`
	Size      float64 `yaml:"size"`
	Type      string  `yaml:"type"`
	TimeDiff  int64   `yaml:"timediff"`
}

type PromeConf struct {
	Switch_on bool `yaml:"switch_on"`
	Port      int  `yaml:"port"`
}

var conf Conf

const (
	G_TestConfName    = "test"
	G_ProductConfName = "product"
	G_TestConfPath    = "../../conf/test/test.yaml"
	G_ProductConfPath = "../../conf/product/product.yaml"
)

func ParseConfInit(confType string) error {
	//打开文件
	var (
		fileF []byte
		err   error
	)
	if confType == G_TestConfName {
		fileF, err = ioutil.ReadFile(G_TestConfPath)
	} else if confType == G_ProductConfName {
		fileF, err = ioutil.ReadFile(G_ProductConfPath)
	} else {
		err = fmt.Errorf("file type not found exception")
	}
	if err != nil {
		log.Printf("ParseConfInit read file error,err:%+v\n", err)
		return err
	}
	err = yaml.Unmarshal(fileF, &conf)
	if err != nil {
		log.Printf("ParseConfInit umarshal error,err:%+v\n", err)
		return err
	}
	return nil
}

func GetParseConfPort() int {
	return conf.Port
}

func GetParseConfPrometheus() PromeConf {
	return conf.Prometheus
}

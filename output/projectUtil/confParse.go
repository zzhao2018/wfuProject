
package projectUtil

import (
    "fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Port  int  `yaml:"port"`
	Prometheus PromeConf `yaml:"prometheus"`
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

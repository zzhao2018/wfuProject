package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

/*
自动代码生成
*/

func main() {
	//初始化cli
	appCli:=cli.NewApp()
	//初始化变量
	var opt GenOption
	appCli.Flags=[]cli.Flag{
		cli.StringFlag{
			Name:        "p",
			Usage:       "prefix of output file",
			Value:       "",
			Destination: &opt.Prefix,
		},
		cli.StringFlag{
			Name:        "o",
			Usage:       "output package path",
			Value:       "./output/",
			Destination: &opt.OutputPath,
		},
		cli.StringFlag{
			Name:        "pro",
			Usage:       "proto file path",
			Value:       "",
			Destination: &opt.ProtoFilePath,
		},
		cli.BoolFlag{
			Name:        "client",
			Usage:       "generator client code",
			Destination: &opt.ClientCode,
		},
		cli.BoolFlag{
			Name:        "server",
			Usage:       "genrator server code",
			Destination: &opt.ServerCode,
		},
	}
	appCli.Action= func(c *cli.Context) error{
		err:=GenMgrInstance.Run(&opt)
		if err!=nil {
			log.Printf("gen mgr error,err:%+v\n",err)
			return err
		}
		return nil
	}
	err:=appCli.Run(os.Args)
	if err!=nil {
		log.Printf("run data error,err:%+v\n",err)
		return
	}
}


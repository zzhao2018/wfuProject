package main

type Generator interface {
	Run(opt *GenOption,metaData *protoMetaData)error
}
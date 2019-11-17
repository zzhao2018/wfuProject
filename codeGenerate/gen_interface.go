package main

type Generator interface {
	Run(opt *GenOption)error
}
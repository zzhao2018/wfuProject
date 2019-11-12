package register

import "time"

/*
  注册中心配置信息
*/
type RegisterOption struct {
	Addr []string
	TimeOut time.Duration
	RegisterPath string
	HeartBeat int64
}

type RegisterOptionFunc func(option *RegisterOption)

func RegisterInitAddr(addrs []string)RegisterOptionFunc{
	return func(option *RegisterOption) {
		option.Addr=addrs
	}
}

func RegisterInitTimeOut(timeout time.Duration)RegisterOptionFunc{
	return func(option *RegisterOption) {
		option.TimeOut=timeout
	}
}

func RegisterInitRegisterPath(registerPath string)RegisterOptionFunc{
	return func(option *RegisterOption) {
		option.RegisterPath=registerPath
	}
}

func RegisterInitHeartBeat(heartbeat int64)RegisterOptionFunc{
	return func(option *RegisterOption) {
		option.HeartBeat=heartbeat
	}
}

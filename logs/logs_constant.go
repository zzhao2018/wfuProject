package logs

import "strings"

//初始化配置
const(
	defaultMessChanSize=200
)

//日志等级
type LogLevel int

const(
	L_Debug=LogLevel(0)
	L_Info=LogLevel(1)
	L_Warn=LogLevel(2)
	L_Error=LogLevel(3)
)

func(l LogLevel)GetLogLevelName()string{
	switch l {
	case L_Debug:
		return "DEBUG"
	case L_Info:
		return "INFO"
	case L_Warn:
		return "WARN"
	case L_Error:
		return "ERROR"
	default:
		return ""
	}
}

func GetLogLevelFromStr(level string)LogLevel{
	level=strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return L_Debug
	case "INFO":
		return L_Info
	case "WARN":
		return L_Warn
	case "ERROR":
		return L_Error
	default:
		return L_Debug
	}
}
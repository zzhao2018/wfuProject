package logs

type LogOutput interface {
	Write(mess *LogMessMeta)error
	Close()
}
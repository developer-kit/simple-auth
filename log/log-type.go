package log

type LogType uint

const (
	LT_CONSOLE = LogType(1)
	LT_FILE    = LogType(2)
	LT_FORMAT  = LogType(3)
	LT_SOCKET  = LogType(4)
	LT_XML     = LogType(5)
)

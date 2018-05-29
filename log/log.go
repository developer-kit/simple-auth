package log

import (
	"fmt"
	"github.com/alecthomas/log4go"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Logger struct {
	logger log4go.Logger
}

func (self *Logger) Init(ltypes ...LogType) {
	self.logger = make(log4go.Logger)
	var logWriter log4go.LogWriter
	for _, ltype := range ltypes {
		if LT_CONSOLE == ltype {
			logWriter = log4go.NewConsoleLogWriter()
		} else if LT_FILE == ltype {
			timeStr := time.Now().Format("2006-01-02_15-04-05")
			fileName := "asr-auth." + timeStr + ".log"
			logWriter = log4go.NewFileLogWriter(fileName, true)
		} else if LT_FORMAT == ltype {
			logWriter = log4go.NewFormatLogWriter(os.Stdout, "[%D %T] [%L] %M")
		} else if LT_SOCKET == ltype {
			logWriter = log4go.NewSocketLogWriter("tcp", "127.0.0.1")
		} else if LT_XML == ltype {
			logWriter = log4go.NewXMLLogWriter("asr-auth.log", true)
		}
		self.logger.AddFilter(strconv.Itoa(int(ltype)), log4go.FINEST, logWriter)
	}
}

func (self Logger) getCurrentLoggerCallFuncName(full bool) string {
	pc, _, lineno, ok := runtime.Caller(3)
	src := ""
	if ok {
		src = runtime.FuncForPC(pc).Name()
		if !full {
			slice := strings.Split(src, "/")
			src = slice[len(slice)-1]
		}
		src = fmt.Sprintf("%s Line:%d", src, lineno)
	}
	return src
}

func (self Logger) Finest(args ...interface{}) {
	self.logger.Finest("(Func: %s) Message:%s", self.getCurrentLoggerCallFuncName(false), args)
}

func (self Logger) Fine(args ...interface{}) {
	self.logger.Fine("(Func: %s) Message:%s", self.getCurrentLoggerCallFuncName(false), args)
}

func (self Logger) Debug(args ...interface{}) {
	self.logger.Debug("(Func: %s) Message:%s", self.getCurrentLoggerCallFuncName(false), args)
}
func (self Logger) Trace(args ...interface{}) {
	self.logger.Trace("(Func: %s) Message:%s", self.getCurrentLoggerCallFuncName(false), args)
}

func (self Logger) Info(args ...interface{}) {
	self.logger.Info("(Func: %s) Message:%s", self.getCurrentLoggerCallFuncName(false), args)
}

func (self Logger) Warn(args ...interface{}) {
	self.logger.Warn("(Func: %s) Message:%s", self.getCurrentLoggerCallFuncName(false), args)
}

func (self Logger) Error(args ...interface{}) {
	self.logger.Error("(Func: %s) Message:%s", self.getCurrentLoggerCallFuncName(false), args)
}

func (self Logger) Critical(args ...interface{}) {
	self.logger.Critical("(Func: %s) Message:%s", self.getCurrentLoggerCallFuncName(false), args)
}

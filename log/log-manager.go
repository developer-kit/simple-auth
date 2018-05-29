package log

import (
	"sync"
)

var logger Logger

type LogManager struct {
}

var m *LogManager
var once sync.Once

func GetInstance() *LogManager {
	once.Do(func() {
		m = &LogManager{}
	})
	return m
}

func (self *LogManager) InitManager() error {
	logger.Init(LT_FORMAT, LT_FILE)
	return nil
}

func Finest(args ...interface{}) {
	logger.Finest(args)
}

func Fine(args ...interface{}) {
	logger.Fine(args)
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Trace(args ...interface{}) {
	logger.Trace(args)
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Warn(args ...interface{}) {
	logger.Warn(args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Critical(args ...interface{}) {
	logger.Critical(args)
}

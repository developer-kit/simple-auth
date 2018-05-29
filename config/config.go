package config

import (
	"errors"
	"github.com/seaofstars-coder/simple-auth/log"
	"github.com/seaofstars-coder/simple-auth/util"
	"github.com/Unknwon/goconfig"
	"time"
)

type ServiceConfig struct {
	conf            *goconfig.ConfigFile
	confName        string
	isLoadSucc      bool
	fileLastModTime int64
	lastCheckTime   int64
}

func (self *ServiceConfig) Init(confFile string) error {
	self.confName = confFile
	var err error = nil
	self.conf, err = goconfig.LoadConfigFile(confFile)
	if nil != err {
		return err
	}
	self.isLoadSucc = true
	self.fileLastModTime, _ = util.GetFileModTime(self.confName)
	self.CheckConfigFile()
	return err
}

func (self ServiceConfig) GetConfig(confName string) (string, error) {
	if !self.isLoadSucc {
		return "", errors.New("GetConfig fail! Is not load success!!!")
	}
	return self.conf.GetValue("", confName)
}

func (self *ServiceConfig) CheckConfigFile() {
	go func() {
		for {
			fileModTime, err := util.GetFileModTime(self.confName)
			if nil != err {
				continue
			}
			if fileModTime != self.fileLastModTime {
				self.conf.Reload()
				self.fileLastModTime = fileModTime
				log.Info(self.confName, " is change, conf reload")
			}
			time.Sleep(time.Second)
		}
	}()
}

package main

import (
	"github.com/seaofstars-coder/simple-auth/auth"
	"github.com/seaofstars-coder/simple-auth/config"
	"github.com/seaofstars-coder/simple-auth/grpc"
	"github.com/seaofstars-coder/simple-auth/http"
	"github.com/seaofstars-coder/simple-auth/log"
	"github.com/seaofstars-coder/simple-auth/redis"
	"github.com/seaofstars-coder/simple-auth/zookeeper"
	"sync"
	"github.com/seaofstars-coder/simple-auth/task"
)

var SyncWG *sync.WaitGroup
func init(){
	SyncWG = new(sync.WaitGroup)
}

func main() {
	initErr := initManager()
	if nil != initErr {
		log.Info(initErr.Error())
		return
	}
	SyncWG.Wait()
}

func initManager() error {
	var err error = nil
	err = log.GetInstance().InitManager()
	if nil != err {
		return err
	}
	err = config.GetInstance().InitManager()
	if nil != err {
		return err
	}
	err = redis.GetInstance().InitManager()
	if nil != err {
		return err
	}
	err = zookeeper.GetInstance().InitManager()
	if nil != err {
		return err
	}
	err = auth.GetInstance().InitManager()
	if nil != err {
		return err
	}
	err = grpc.GetInstance().InitManager(SyncWG)
	if nil != err {
		return err
	}
	err = http.GetInstance().InitManager(SyncWG)
	if nil != err {
		return err
	}
	err = task.GetInstance().InitManager()
	if nil != err {
		return err
	}
	return err
}

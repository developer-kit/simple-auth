package task

import (
	"errors"
	"github.com/seaofstars-coder/simple-auth/log"
	"sync"
)

type TaskManager struct {
	taskMap TaskMap
	mutex   sync.Mutex
}

var m *TaskManager
var once sync.Once

func GetInstance() *TaskManager {
	once.Do(func() {
		m = &TaskManager{}
	})
	return m
}

func (self *TaskManager) InitManager() error {
	self.taskMap = make(TaskMap)
	self.initTaskMap()
	return nil
}

func (self *TaskManager) initTaskMap() {

}

func (self *TaskManager) RegisterTask(ttype TaskType, te TaskEntiy) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	var err error = nil
	self.taskMap[ttype], err = TaskCreate(ttype)
	if nil != err {
		log.Error("RegisterTask fail! task type is %v\n", ttype)
		return
	}
	self.taskMap[ttype].Init(te.param, te.f)
}

func (self *TaskManager) UnregisterTask(ttype TaskType) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	task, err := self.getTaskByType(ttype)
	if nil == err {
		log.Error("UnregisterTask fail!, has not task exist !!! task type is %v\n", ttype)
		return
	}
	task.Terminate()
	task = nil
}

func (self *TaskManager) getTaskByType(ttype TaskType) (Task, error) {
	if nil == self.taskMap[ttype] {
		log.Error("UnregisterTask fail!, has not task exist !!! task type is %v\n", ttype)
		return nil, errors.New("getTaskByType fail!, has not task exist !!!")
	}
	return self.taskMap[ttype], nil
}

func (self *TaskManager) SuspendTask(ttype TaskType) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	task, err := self.getTaskByType(ttype)
	if nil == err {
		log.Error("SuspendTask fail!, has not task exist !!! task type is %v\n", ttype)
		return
	}
	task.Suspend()
}

func (self *TaskManager) SleepTask(ttype TaskType, st int) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	task, err := self.getTaskByType(ttype)
	if nil == err {
		log.Error("SuspendTask fail!, has not task exist !!! task type is %v\n", ttype)
		return
	}
	task.Sleep(st)
}

func (self *TaskManager) ResumeTask(ttype TaskType) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	task, err := self.getTaskByType(ttype)
	if nil == err {
		log.Error("ResumeTask fail!, has not task exist !!! task type is %v\n", ttype)
		return
	}
	if TST_SLEEPED == task.GetStatus() || TST_SUSPENDED == task.GetStatus() {
		task.Resume()
	}
}

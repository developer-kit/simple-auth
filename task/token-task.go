package task

import (
	"github.com/seaofstars-coder/simple-auth/util"
)

type TokenTask struct {
	param       TaskParam
	f           TaskFunc
	lastRunTime int64
	status      TaskStatusType
}

func (self *TokenTask) Init(tp TaskParam, tf TaskFunc) error {
	self.param = tp
	self.f = tf
	self.start()
	return nil
}

func (self *TokenTask) GetStatus() TaskStatusType {
	return self.status
}

func (self *TokenTask) Terminate() error {
	self.status = TST_STOPED
	return nil
}

func (self *TokenTask) Suspend() error {
	self.status = TST_SLEEPED
	return nil
}

func (self *TokenTask) Resume() error {
	self.status = TST_RUNNING
	return nil
}

func (self *TokenTask) Sleep(time int) error {
	return nil
}

func (self *TokenTask) start() error {
	go func() {
		self.status = TST_RUNNING
		for {
			if TST_STOPED == self.status {
				break
			}
			nowt := util.GetCurrentSeconds()
			self.update(nowt)
		}
	}()
	return nil
}

func (self *TokenTask) update(nowt int64) error {
	if TST_RUNNING == self.status {
		if util.CheckTimeLargeOrEqual(nowt, self.lastRunTime, self.param.interval) {
			self.lastRunTime = nowt
			self.f()
		}
	}
	return nil
}

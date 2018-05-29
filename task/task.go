package task

import "errors"

type Task interface {
	Init(tp TaskParam, tf TaskFunc) error
	GetStatus() TaskStatusType
	Terminate() error
	Suspend() error
	Resume() error
	Sleep(time int) error
	start() error
	update(nowt int64) error
}

func TaskCreate(ttype TaskType) (Task, error) {
	if TT_AUTH_TASK_CHECK == ttype {
		return new(TokenTask), nil
	}
	return nil, errors.New("TaskCreate fail! Has not task type!!!")
}

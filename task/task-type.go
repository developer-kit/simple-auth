package task

type TaskType uint

const (
	TT_AUTH_TASK_CHECK = TaskType(1)
)

type TaskStatusType uint

const (
	TST_RUNNING   = TaskStatusType(1)
	TST_SLEEPED   = TaskStatusType(2)
	TST_SUSPENDED = TaskStatusType(3)
	TST_STOPED    = TaskStatusType(4)
)

type TaskParam struct {
	interval int64
}

type TaskFunc func()

type TaskEntiy struct {
	param TaskParam
	f     TaskFunc
}

type TaskMap map[TaskType]Task

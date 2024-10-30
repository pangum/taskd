package kernel

import (
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/model"
)

var _ task.Task = (*Task)(nil)

type Task struct {
	task *model.Task
}

func NewTask(task *model.Task) *Task {
	return &Task{
		task: task,
	}
}

func (t *Task) Id() uint64 {
	return t.task.Id
}

func (t *Task) Target() uint64 {
	return t.task.Target
}

func (t *Task) Type() task.Type {
	return t.task.Type
}

func (t *Task) Subtype() task.Type {
	return t.task.Subtype
}

func (t *Task) Retries() uint32 {
	return t.task.Retries
}

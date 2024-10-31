package kernel

import (
	"time"

	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/model"
)

var _ task.Task = (*Task)(nil)

type Task struct {
	task *model.Tasker
}

func NewTask(task *model.Tasker) *Task {
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

func (t *Task) Elapsed() time.Duration {
	return t.task.Elapsed
}

func (t *Task) Data() map[string]any {
	return t.task.Data
}

func (t *Task) Next() time.Time {
	return t.task.Next
}

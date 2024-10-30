package kernel

import (
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/model"
	"github.com/pangum/taskd/internal/internal/repository"
)

var _ task.Task = (*Task)(nil)

type Task struct {
	task       *model.Task
	repository repository.Schedule
}

func NewTask(task *model.Task, repository repository.Schedule) *Task {
	return &Task{
		task:       task,
		repository: repository,
	}
}

func (t *Task) Id() uint64 {
	return t.task.Id
}

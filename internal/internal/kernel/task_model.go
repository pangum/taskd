package kernel

import (
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/model"
	"github.com/pangum/taskd/internal/internal/repository"
)

var _ task.Task = (*TaskModel)(nil)

type TaskModel struct {
	task       *model.Task
	repository repository.Schedule
}

func NewTaskModel(task *model.Task, repository repository.Schedule) *TaskModel {
	return &TaskModel{
		task:       task,
		repository: repository,
	}
}

func (t *TaskModel) Id() uint64 {
	return t.task.Id
}

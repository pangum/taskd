package core

import (
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/model"
)

type Schedule interface {
	Add(schedule task.Schedule) error

	Get(schedule *model.Schedule, columns ...string) (bool, error)

	Update(schedule *model.Schedule, columns ...string) (int64, error)

	Delete(schedule *model.Schedule) (int64, error)
}

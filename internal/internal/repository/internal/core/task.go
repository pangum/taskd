package core

import (
	"github.com/pangum/taskd/internal/internal/model"
)

type Task interface {
	Add(task *model.Task) (int64, error)

	Get(task *model.Task, columns ...string) (bool, error)

	GetsRunnable(times uint32, excludes ...*model.Task) (*[]*model.Task, error)

	Update(task *model.Task, columns ...string) (int64, error)

	Delete(task *model.Task) (int64, error)
}

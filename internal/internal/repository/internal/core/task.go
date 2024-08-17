package core

import (
	"time"

	"github.com/pangum/taskd/internal/internal/model"
)

type Task interface {
	Add(task *model.Task) (int64, error)

	Get(task *model.Task, columns ...string) (bool, error)

	GetsRunnable(count int, times uint32, maximum time.Duration, excludes ...*model.Task) (*[]*model.Task, error)

	Update(task *model.Task, columns ...string) (int64, error)

	Delete(task *model.Task) (int64, error)
}

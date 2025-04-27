package core

import (
	model2 "github.com/harluo/taskd/internal/internal/internal/model"
)

type Task interface {
	Add(task *model2.Task) (int64, error)

	Get(task *model2.Task, columns ...string) (bool, error)

	GetsRunnable(excludes ...*model2.Task) (*[]*model2.Tasker, error)

	Update(task *model2.Task, columns ...string) (int64, error)

	Archive(task *model2.Task) (int64, error)

	Delete(task *model2.Task) (int64, error)
}

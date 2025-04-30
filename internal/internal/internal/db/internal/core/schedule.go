package core

import (
	model2 "github.com/harluo/taskd/internal/internal/internal/model"
)

type Schedule interface {
	Add(*model2.Runtime, ...*model2.Runtime) (*[]*model2.Tasker, error)

	Get(*model2.Schedule, ...string) (bool, error)

	Update(*model2.Schedule, ...string) (int64, error)

	Delete(*model2.Schedule) (int64, error)
}

package core

import (
	"time"

	"github.com/pangum/taskd/internal/internal/model"
)

type Schedule interface {
	Add(*model.Schedule, time.Time) (*model.Tasker, error)

	Get(*model.Schedule, ...string) (bool, error)

	Update(*model.Schedule, ...string) (int64, error)

	Delete(*model.Schedule) (int64, error)
}

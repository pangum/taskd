package internal

import (
	"time"

	"github.com/goexl/log"
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/schedule"
)

var _ task.Tasker = (*Tasker)(nil)

type Tasker struct {
	runnable schedule.Runnable
	logger   log.Logger
}

func NewTasker(runnable schedule.Runnable, logger log.Logger) task.Tasker {
	return &Tasker{
		runnable: runnable,
		logger:   logger,
	}
}

func (t Tasker) Add(scheduling task) error {
	// TODO implement me
	panic("implement me")
}

func (t Tasker) Remove(scheduling interface{}) error {
	// TODO implement me
	panic("implement me")
}

func (t Tasker) Running(id uint64, status interface{}, retries uint32) error {
	// TODO implement me
	panic("implement me")
}

func (t Tasker) Update(id uint64, status interface{}, runtime time.Time) error {
	// TODO implement me
	panic("implement me")
}

func (t Tasker) Next(id uint64) error {
	// TODO implement me
	panic("implement me")
}

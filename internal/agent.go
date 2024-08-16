package internal

import (
	"github.com/goexl/log"
	"github.com/goexl/schedule"
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/config"
	"github.com/pangum/taskd/internal/internal/repository"
)

var _ task.Agent = (*Agent)(nil)

type Agent struct {
	task      repository.Task
	scheduler *schedule.Scheduler
	config    *config.Retry
	logger    log.Logger
}

func NewAgent(task repository.Task) *Agent {
	return &Agent{
		task: task,
	}
}

func (a Agent) Start(processor task.Processor) (err error) {
	return
}

func (a Agent) Add(scheduling task.Scheduling) (err error) {
	return
}

func (a Agent) Remove(scheduling task.Scheduling) (err error) {
	return
}

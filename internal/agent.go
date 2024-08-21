package internal

import (
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/schedule"
)

var _ task.Agent = (*Agent)(nil)

type Agent struct {
	runnable schedule.Runnable
	logger   log.Logger
}

func NewAgent(runnable schedule.Runnable, logger log.Logger) *Agent {
	return &Agent{
		runnable: runnable,
		logger:   logger,
	}
}

func (a Agent) Start(processor task.Processor) (err error) {
	if sre := a.runnable.Start(processor); nil != sre {
		err = sre
		a.logger.Error("启动任务处理出错", field.Error(err))
	} else {
		a.logger.Info("启动任务处理成功", field.New("processor", processor))
	}

	return
}

func (a Agent) Add(scheduling task.Scheduling) (err error) {
	return
}

func (a Agent) Remove(scheduling task.Scheduling) (err error) {
	return
}

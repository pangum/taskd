package internal

import (
	"context"

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

func NewAgent(runnable schedule.Runnable, logger log.Logger) task.Agent {
	return &Agent{
		runnable: runnable,
		logger:   logger,
	}
}

func (a Agent) Start(_ context.Context, processor task.Processor) (err error) {
	if sre := a.runnable.Start(processor); nil != sre {
		err = sre
		a.logger.Error("启动任务处理出错", field.Error(err))
	} else {
		a.logger.Info("启动任务处理成功", field.New("processor", processor))
	}

	return
}

func (a Agent) Add(_ task.Scheduling) (err error) {
	return
}

func (a Agent) Remove(_ task.Scheduling) (err error) {
	return
}

func (a Agent) Stop(_ context.Context) (err error) {
	if sre := a.runnable.Stop(); nil != sre {
		err = sre
		a.logger.Error("停止任务处理出错", field.Error(err))
	} else {
		a.logger.Info("停止任务处理成功", field.New("scheduler", "runnable"))
	}

	return
}

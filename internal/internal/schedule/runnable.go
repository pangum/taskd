package schedule

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/schedule"
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/config"
	"github.com/pangum/taskd/internal/internal/model"
	"github.com/pangum/taskd/internal/internal/repository"
	"github.com/pangum/taskd/internal/internal/schedule/internal/core"
	"github.com/pangum/taskd/internal/internal/schedule/internal/executor"
	"github.com/pangum/taskd/internal/internal/schedule/internal/get"
)

type Runnable struct {
	repository repository.Task
	scheduler  *schedule.Scheduler
	executor   *executor.Default
	processor  task.Processor
	config     *config.Retry
	running    *core.Running
	logger     log.Logger
	id         string
}

func newRunnable(get get.Runnable) *Runnable {
	return &Runnable{
		repository: get.Repository,
		scheduler:  get.Scheduler,
		executor:   get.Executor,
		config:     get.Config,
		running:    get.Running,
		logger:     get.Logger,
	}
}

func (r *Runnable) Start(processor task.Processor) (err error) {
	name := "任务执行器"
	fields := gox.Fields[any]{
		field.New("name", name),
	}
	if id, ae := r.scheduler.Add(r).Duration(r.config.Interval).Name(name).Build().Apply(); nil != ae {
		err = ae
		r.logger.Error("添加任务出错", fields.Add(field.Error(ae))...)
	} else {
		r.processor = processor
		r.id = id
		r.logger.Info("添加任务成功", fields.Add(field.New("id", id))...)
	}

	return
}

func (r *Runnable) Stop() {
	r.scheduler.Remove().Id(r.id).Build().Apply()
}

func (r *Runnable) Run() (err error) {
	count := r.config.Count
	times := r.config.Times
	maximum := r.config.Maximum
	if tasks, re := r.repository.GetsRunnable(count, times, maximum, r.running.Tasks()...); nil != re {
		err = re
	} else if 0 != len(*tasks) {
		err = r.add(tasks)
	}

	return
}

func (r *Runnable) add(tasks *[]*model.Task) (err error) {
	for _, _task := range *tasks {
		builder := r.scheduler.Add(r.executor.Clone(_task, r.processor))
		// 主动设置标识，确保能唯一标识一个任务
		// 必须保证一个任务同一时间只有一次调度
		builder.Id(_task)
		builder.Fixed(_task.Next)

		fields := gox.Fields[any]{
			field.New("_task", _task),
		}
		if id, de := builder.Build().Apply(); nil != de {
			r.logger.Info("添加任务出错", fields.Add(field.Error(de))...)
		} else {
			r.running.Add(_task)
			r.logger.Debug("添加任务成功", fields.Add(field.New("id", id))...)
		}
	}

	return
}

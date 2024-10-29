package core

import (
	"context"
	"sync"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/schedule"
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/get"
	"github.com/pangum/taskd/internal/internal/repository"
)

type Tasker struct {
	schedule  repository.Schedule
	task      repository.Task
	scheduler *schedule.Scheduler

	runnable *sync.Map
	logger   log.Logger
	id       string
	times    uint32
}

func newTasker(get get.Runnable) *Tasker {
	return &Tasker{
		task:      get.Repository,
		scheduler: get.Scheduler,
		logger:    get.Logger,
	}
}

func (t *Tasker) Start(_ context.Context) (err error) {
	name := "任务执行器"
	fields := gox.Fields[any]{
		field.New("name", name),
	}
	if id, ae := t.scheduler.Add(t).Duration(3 * time.Second).Name(name).Build().Apply(); nil != ae {
		err = ae
		t.logger.Error("添加任务出错", fields.Add(field.Error(ae))...)
	} else {
		t.id = id
		t.logger.Info("添加任务成功", fields.Add(field.New("id", id))...)
	}

	return
}

func (t *Tasker) Stop(_ context.Context) (err error) {
	t.scheduler.Remove().Id(t.id).Build().Apply()
	err = t.scheduler.Stop()

	return
}

func (t *Tasker) Add(schedule task.Schedule) (err error) {
	if ae := t.schedule.Add(schedule); nil != ae {
		err = ae
	} else if schedule.Next().Before(time.Now()) {
		t.runnable.Store(schedule.Id(), schedule) // TODO 放进准备运行的列表
	}

	return
}

func (t *Tasker) Remove(scheduling task.Schedule) (err error) {
	return
}

func (t *Tasker) Running(id uint64, status task.Status, retries uint32) (err error) {
	return
}

func (t *Tasker) Update(id uint64, status task.Status, runtime time.Time) (err error) {
	return
}

func (t *Tasker) Next(id uint64) (err error) {
	return
}

func (t *Tasker) Pop() (runnable task.Task, exists bool) {
	t.runnable.Range(func(key, value any) (next bool) {
		runnable = value.(task.Task)
		next = false

		return
	})

	return
}

func (t *Tasker) Run() (err error) {
	if tasks, re := t.task.GetsRunnable(t.times); nil != re {
		err = re
	} else if 0 != len(*tasks) {
		for _, _task := range *tasks {
			t.runnable.Store(_task.Id, _task)
		}
	}

	return
}

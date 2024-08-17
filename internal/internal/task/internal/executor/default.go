package executor

import (
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/structer"
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/config"
	"github.com/pangum/taskd/internal/internal/internal/column"
	"github.com/pangum/taskd/internal/internal/model"
	"github.com/pangum/taskd/internal/internal/repository"
	"github.com/pangum/taskd/internal/internal/task/internal/core"
	"github.com/pangum/taskd/internal/internal/task/internal/executor/internal"
	"github.com/pangum/taskd/internal/internal/task/internal/executor/internal/get"
)

type Default struct {
	repository repository.Task
	config     *config.Retry
	running    *core.Running
	logger     log.Logger

	processor task.Processor // 执行器
	task      *model.Task    // 用于保存每次任务的执行原始数据
	executor  task.Executor  // 用于清理数据时计算下一次执行时间
}

func newDefault(get get.Default) *Default {
	return &Default{
		repository: get.Repository,
		config:     get.Config,
		running:    get.Running,
		logger:     get.Logger,
	}
}

func (d *Default) Clone(task *model.Task, processor task.Processor) (cloned *Default) {
	cloned = new(Default)
	if ce := structer.Copy().From(d).To(cloned).Build().Apply(); nil == ce {
		cloned.task = task
		cloned.processor = processor
	} else {
		d.logger.Warn("创建执行器出错", field.New("task", task))
	}

	return
}

func (d *Default) Run() (err error) {
	defer func() {
		err = d.cleanup(&err)
	}()

	if exists, ge := d.repository.Get(d.task); nil != ge {
		err = ge
	} else if !exists {
		d.logger.Warn("要被执行的任务已不存在", field.New("task", d.task))
	} else if re := d.updateRunning(); nil != re {
		err = re
	} else {
		err = d.run()
	}

	return
}

func (d *Default) run() (err error) {
	tasking := internal.NewTasking(d.task.Id, d.task.Target, d.task.Data, d.task.Times)
	if executor, pe := d.processor.Process(tasking); nil != pe {
		err = pe
	} else {
		d.executor = executor
		err = d.executor.Execute()
	}

	return
}

func (d *Default) updateRunning() (err error) {
	updated := new(model.Task)
	updated.Id = d.task.Id
	updated.Status = gox.Ift[task.Status](0 == d.task.Times, task.StatusRunning, task.StatusRetrying)
	updated.Times = d.task.Times + 1
	err = d.update(updated)

	return
}

func (d *Default) cleanup(result *error) (err error) {
	if nil == *result { // 执行成功，更新下一次执行时间
		d.running.Remove(d.task)
		err = d.next()
	} else if d.task.Times >= d.config.Times {
		d.running.Remove(d.task)
		d.logger.Warn("已达到最大重试次数，任务不再被执行", field.New("task", d.task), field.New("time", d.task.Times))
		err = d.next()
	} else { // 以二的幂为基数重试
		updated := new(model.Task)
		updated.Id = d.task.Id
		updated.Status = task.StatusFailed
		// 确定下一次重试的时间，计算规则是，以二的幂为基数重试
		updated.Next = time.Now().Add(15 * time.Second * 2 << d.task.Times)
		err = d.update(updated)
	}

	return
}

func (d *Default) update(task *model.Task, columns ...string) (err error) {
	task.Version = d.task.Version
	if _, ue := d.repository.Update(task, columns...); nil != ue {
		err = ue
		d.logger.Warn("更新任务出错", field.New("task", task), field.Error(ue))
	} else {
		d.task.Version = task.Version // 确保下一次更新乐观锁必须一致
		d.logger.Debug("更新任务成功", field.New("task", task))
	}

	return
}

func (d *Default) next() (err error) {
	if d.executor.Recyclable() { // 需要循环执行，更新下一次执行时间，等待被调度
		err = d.updateNext()
	} else { // 不需要被循环执行，清理数据
		deleted := new(model.Task)
		deleted.Id = d.task.Id
		_, err = d.repository.Delete(deleted)
	}

	return
}

func (d *Default) updateNext() (err error) {
	updated := new(model.Task)
	updated.Id = d.task.Id
	updated.Times = 0
	updated.Status = task.StatusSuccess
	updated.Next = d.executor.Next()
	err = d.update(updated, column.Times.String())

	return
}

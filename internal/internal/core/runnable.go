package core

import (
	"github.com/goexl/collection"
	"github.com/goexl/guc/collection/queue"
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/kernel"
	"github.com/pangum/taskd/internal/internal/model"
)

type Runnable struct {
	tasks collection.Queue[task.Task]
}

func newRunnable() *Runnable {
	return &Runnable{
		tasks: queue.NewBlocking[task.Task]().Build(),
	}
}

func (r *Runnable) Put(required *model.Tasker, others ...*model.Tasker) {
	for _, _task := range append([]*model.Tasker{required}, others...) {
		r.tasks.Enqueue(kernel.NewTask(_task))
	}
}

func (r *Runnable) Tasks() (tasks []task.Task) {
	return r.tasks.Dequeue()
}

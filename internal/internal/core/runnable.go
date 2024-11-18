package core

import (
	"time"

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

func (r *Runnable) Put(required *model.Tasker, optionals ...*model.Tasker) {
	for _, tasker := range append([]*model.Tasker{required}, optionals...) {
		if tasker.Next.Before(time.Now()) {
			r.tasks.Enqueue(kernel.NewTask(tasker))
		}
	}
}

func (r *Runnable) Tasks() (tasks []task.Task) {
	return r.tasks.Dequeue()
}

package core

import (
	"time"

	"github.com/goexl/container"
	"github.com/goexl/container/queue"
	"github.com/goexl/task"
	"github.com/harluo/taskd/internal/internal/kernel"
	"github.com/harluo/taskd/internal/internal/model"
)

type Runnable struct {
	tasks container.Queue[task.Task]
}

func newRunnable() *Runnable {
	return &Runnable{
		tasks: queue.New[task.Task]().Blocking().Build(),
	}
}

func (r *Runnable) Put(required *model.Tasker, optionals ...*model.Tasker) {
	for _, tasker := range append([]*model.Tasker{required}, optionals...) {
		if tasker.Next.Before(time.Now()) {
			r.tasks.Enqueue(kernel.NewTask(tasker))
		}
	}
}

func (r *Runnable) Task() task.Task {
	return r.tasks.Dequeue()
}

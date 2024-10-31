package core

import (
	"github.com/goexl/task"
	"github.com/pangum/taskd/internal/internal/kernel"
	"github.com/pangum/taskd/internal/internal/model"

	"sync"
)

type Runnable struct {
	tasks *sync.Map
}

func NewRunnable() *Runnable {
	return &Runnable{
		tasks: new(sync.Map),
	}
}

func (r *Runnable) Put(required *model.Tasker, others ...*model.Tasker) {
	for _, _task := range append([]*model.Tasker{required}, others...) {
		r.tasks.Store(_task.Id, kernel.NewTask(_task))
	}
}

func (r *Runnable) Tasks() (tasks []task.Task) {
	r.tasks.Range(func(key, value any) (next bool) {
		tasks = append(tasks, value.(task.Task))
		r.tasks.Delete(key)
		next = true

		return
	})

	return
}

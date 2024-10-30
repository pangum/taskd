package core

import (
	"github.com/goexl/task"
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

func (r *Runnable) Put(required *model.Task, others ...*model.Task) {
	for _, _task := range append([]*model.Task{required}, others...) {

	}
}

func (r *Runnable) Tasks() (tasks []task.Task) {
	r.tasks.Range(func(key, value interface{}) (next bool) {
		tasks = append(tasks, value.(task.Task))
		next = true

		return
	})

	return
}

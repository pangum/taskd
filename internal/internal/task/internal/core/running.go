package core

import (
	"sync"

	"github.com/pangum/taskd/internal/internal/model"
)

type Running struct {
	cache    *sync.Map
	tasks    []*model.Task
	modified bool
}

func newRunning() *Running {
	return &Running{
		cache: new(sync.Map),
	}
}

func (r *Running) Add(task *model.Task) {
	r.cache.Store(task.Id, task)
}

func (r *Running) Remove(task *model.Task) {
	r.cache.Delete(task.Id)
}

func (r *Running) Tasks() (tasks []*model.Task) {
	if !r.modified {
		tasks = r.tasks
	} else {
		r.tasks = make([]*model.Task, 0)
		r.cache.Range(func(key, value any) (next bool) {
			r.tasks = append(r.tasks, value.(*model.Task))
			next = true

			return
		})
		tasks = r.tasks
	}

	return
}

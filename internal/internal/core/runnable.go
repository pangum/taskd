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

}

func (r *Runnable) Tasks() (tasks []task.Task) {
	return
}

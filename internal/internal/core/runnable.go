package core

import (
	"github.com/goexl/task"
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

func (r *Runnable) Put(required task.Task, others ...task.Task) {

}

func (r *Runnable) Tasks() (tasks []task.Task) {
	return
}

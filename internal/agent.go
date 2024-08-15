package internal

import (
	"github.com/pangum/taskd/internal/internal/repository"
)

type Agent struct {
	task repository.Task
}

func NewAgent(task repository.Task) *Agent {
	return &Agent{
		task: task,
	}
}

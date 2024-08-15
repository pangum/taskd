package repository

import (
	"github.com/pangum/taskd/internal/internal/repository/internal/core"
	"github.com/pangum/taskd/internal/internal/repository/internal/get"
	"github.com/pangum/taskd/internal/internal/repository/internal/mysql"
)

// Task 任务
type Task = core.Task

func newTask(database get.Database) Task {
	return mysql.NewTask(database)
}

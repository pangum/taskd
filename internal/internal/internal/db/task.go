package db

import (
	"github.com/harluo/taskd/internal/internal/internal/db/internal/core"
	"github.com/harluo/taskd/internal/internal/internal/db/internal/get"
	"github.com/harluo/taskd/internal/internal/internal/db/internal/mysql"
)

// Task 任务
type Task = core.Task

func newTask(tx get.Transaction) Task {
	return mysql.NewTask(tx)
}

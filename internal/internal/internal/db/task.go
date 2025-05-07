package db

import (
	"github.com/harluo/taskd/internal/internal/internal/db/internal/core"
	"github.com/harluo/taskd/internal/internal/internal/db/internal/get"
	"github.com/harluo/taskd/internal/internal/internal/db/internal/sql"
)

// Task 任务
type Task = core.Task

func newTask(tx get.Tx) Task {
	return sql.NewTask(tx)
}

package db

import (
	"github.com/harluo/taskd/internal/internal/internal/db/internal/core"
	"github.com/harluo/taskd/internal/internal/internal/db/internal/get"
	"github.com/harluo/taskd/internal/internal/internal/db/internal/mysql"
)

// Schedule 计划
type Schedule = core.Schedule

func newSchedule(tx get.Transaction) Schedule {
	return mysql.NewSchedule(tx)
}

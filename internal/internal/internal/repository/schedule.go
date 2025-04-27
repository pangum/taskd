package repository

import (
	"github.com/harluo/taskd/internal/internal/internal/repository/internal/core"
	"github.com/harluo/taskd/internal/internal/internal/repository/internal/get"
	"github.com/harluo/taskd/internal/internal/internal/repository/internal/mysql"
)

// Schedule 计划
type Schedule = core.Schedule

func newSchedule(tx get.Transaction) Schedule {
	return mysql.NewSchedule(tx)
}

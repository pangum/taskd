package repository

import (
	"github.com/pangum/taskd/internal/internal/repository/internal/core"
	"github.com/pangum/taskd/internal/internal/repository/internal/get"
	"github.com/pangum/taskd/internal/internal/repository/internal/mysql"
)

// Schedule 计划
type Schedule = core.Schedule

func newSchedule(tx get.Transaction) Schedule {
	return mysql.NewSchedule(tx)
}

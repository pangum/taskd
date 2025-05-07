package db

import (
	"github.com/harluo/taskd/internal/internal/internal/db/internal/core"
	"github.com/harluo/taskd/internal/internal/internal/db/internal/get"
	"github.com/harluo/taskd/internal/internal/internal/db/internal/sql"
)

// Schedule 计划
type Schedule = core.Schedule

func newSchedule(tx get.Tx) Schedule {
	return sql.NewSchedule(tx)
}

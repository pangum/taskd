package db

import (
	"github.com/harluo/di"
	_ "github.com/harluo/taskd/internal/internal/internal/db/internal/migrate"
)

func init() {
	di.New().Instance().Put(
		newSchedule,
		newTask,
	).Build().Apply()
}

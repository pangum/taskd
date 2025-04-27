package repository

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newSchedule,
		newTask,
	).Build().Apply()
}

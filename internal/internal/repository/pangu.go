package repository

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newSchedule,
		newTask,
	).Build().Apply()
}

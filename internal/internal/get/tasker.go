package get

import (
	"github.com/goexl/log"
	"github.com/pangum/pangu"
	"github.com/pangum/schedule"
	"github.com/pangum/taskd/internal/internal/core"
	"github.com/pangum/taskd/internal/internal/repository"
)

type Tasker struct {
	pangu.Get

	Schedule repository.Schedule
	Task     repository.Task

	Runnable  *core.Runnable
	Scheduler *schedule.Scheduler
	Logger    log.Logger
}

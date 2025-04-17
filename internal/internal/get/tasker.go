package get

import (
	"github.com/goexl/log"
	"github.com/harluo/di"
	"github.com/harluo/taskd/internal/internal/core"
	"github.com/harluo/taskd/internal/internal/repository"
	"github.com/pangum/schedule"
)

type Tasker struct {
	di.Get

	Schedule repository.Schedule
	Task     repository.Task

	Runnable  *core.Runnable
	Scheduler *schedule.Scheduler
	Logger    log.Logger
}

package get

import (
	"github.com/goexl/log"
	"github.com/harluo/di"
	"github.com/harluo/schedule"
	"github.com/harluo/taskd/internal/internal/core"
	"github.com/harluo/taskd/internal/internal/repository"
)

type Tasker struct {
	di.Get

	Schedule repository.Schedule
	Task     repository.Task

	Runnable  *core.Runnable
	Scheduler *schedule.Scheduler
	Logger    log.Logger
}

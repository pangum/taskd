package get

import (
	"github.com/goexl/log"
	"github.com/harluo/di"
	"github.com/harluo/schedule"
	"github.com/harluo/taskd/internal/internal/internal/core"
	repository2 "github.com/harluo/taskd/internal/internal/internal/repository"
)

type Tasker struct {
	di.Get

	Schedule repository2.Schedule
	Task     repository2.Task

	Runnable  *core.Runnable
	Scheduler *schedule.Scheduler
	Logger    log.Logger
}

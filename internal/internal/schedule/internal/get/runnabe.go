package get

import (
	"github.com/goexl/log"
	"github.com/goexl/schedule"
	"github.com/pangum/pangu"
	"github.com/pangum/taskd/internal/internal/config"
	"github.com/pangum/taskd/internal/internal/repository"
	"github.com/pangum/taskd/internal/internal/schedule/internal/core"
	"github.com/pangum/taskd/internal/internal/schedule/internal/executor"
)

type Runnable struct {
	pangu.Get

	Repository repository.Task
	Scheduler  *schedule.Scheduler
	Executor   *executor.Default
	Config     *config.Retry
	Running    *core.Running
	Logger     log.Logger
}

package get

import (
	"github.com/goexl/log"
	"github.com/goexl/schedule"
	"github.com/pangum/pangu"
	"github.com/pangum/taskd/internal/internal/repository"
)

type Runnable struct {
	pangu.Get

	Repository repository.Task
	Scheduler  *schedule.Scheduler
	Logger     log.Logger
}

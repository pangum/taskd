package get

import (
	"github.com/goexl/log"
	"github.com/pangum/pangu"
	"github.com/pangum/taskd/internal/internal/config"
	"github.com/pangum/taskd/internal/internal/repository"
	"github.com/pangum/taskd/internal/internal/task/internal/core"
)

type Default struct {
	pangu.Get

	Repository repository.Task
	Config     *config.Retry
	Running    *core.Running
	Logger     log.Logger
}

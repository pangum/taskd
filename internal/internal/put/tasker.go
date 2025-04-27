package put

import (
	"github.com/harluo/di"
	"github.com/harluo/taskd/internal/internal/service"
)

type Tasker struct {
	di.Put

	Tasker *service.Tasker `name:"harluo.taskd"`
}

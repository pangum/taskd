package put

import (
	"github.com/goexl/task"
	"github.com/harluo/di"
)

type Tasker struct {
	di.Put

	Tasker task.Tasker `name:"harluo.taskd"`
}

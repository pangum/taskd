package taskd

import (
	_ "github.com/pangum/taskd/internal/service"

	"github.com/goexl/task"
)

// Agent 方便外部引用 task.Agent
type Agent = task.Agent

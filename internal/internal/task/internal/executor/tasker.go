package executor

import (
	"time"
)

type Tasker interface {
	// Exec 执行任务逻辑
	Exec() error

	// Recyclable 是否继续执行
	Recyclable() bool

	// Next 下次执行时间
	Next() time.Time

	// Error 执行错误
	Error() error
}

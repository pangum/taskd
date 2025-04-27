package model

import (
	"time"
)

// Runtime 运行时
type Runtime struct {
	Schedule

	Next time.Time
}

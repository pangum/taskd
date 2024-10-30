package model

import (
	"fmt"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/model"
	"github.com/goexl/task"
)

type Task struct {
	model.Base       `xorm:"extends"`
	model.Optimistic `xorm:"extends"` // 使用乐观锁，优化重试时的性能

	// 计划
	Schedule uint64 `xorm:"bigint notnull index(target) default(0) comment(计划)" json:"schedule,omitempty"`
	// 开始执行时间
	// nolint:lll
	Start time.Time `xorm:"datetime notnull index(next) default(CURRENT_TIMESTAMP) comment(开始时间)" json:"start,omitempty"`
	// 下一次重试时间
	// nolint:lll
	Next time.Time `xorm:"datetime notnull index(next) default(CURRENT_TIMESTAMP) comment(一下次执行时间)" json:"next,omitempty"`
	// 结束时间
	// nolint:lll
	Stop time.Time `xorm:"datetime notnull index(next) default(CURRENT_TIMESTAMP) comment(结束时间)" json:"stop,omitempty"`
	// 重试次数
	Retries uint32 `xorm:"int notnull default(0) comment(重试次数)" json:"retries,omitempty"`
	// 状态
	// nolint:lll
	Status task.Status `xorm:"tinyint notnull index(next) default(0) comment(状态，分别是：1、已创建；2、执行中；3、重试中；10、失败；20、成功)" json:"status,omitempty"`

	// 以下字段用于取值而不能存到数据中
	Scheduler Schedule `xorm:"extends <-"`
}

func (*Task) TableComment() string {
	return "任务"
}

func (t *Task) TaskId() (id string) {
	switch {
	case 0 != t.Id:
		id = gox.ToString(t.Id)
	default:
		id = fmt.Sprintf("%p", t)
	}

	return
}

package model

import (
	"fmt"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/id"
	"github.com/goexl/model"
	"github.com/goexl/task"
	"github.com/pangum/pangu"
)

type Task struct {
	model.Base `xorm:"extends"`

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
	Times uint32 `xorm:"int notnull default(0) comment(重试次数)" json:"times,omitempty"`
	// 状态
	// nolint:lll
	Status task.Status `xorm:"tinyint notnull index(next) default(0) comment(状态，分别是：1、已创建；2、执行中；3、重试中；10、失败；20、成功)" json:"status,omitempty"`
}

func (t *Task) BeforeInsert() {
	if 0 == t.Id {
		pangu.New().Get().Dependency().Get(t.setId).Build().Build().Apply()
	}
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

func (t *Task) setId(generator id.Generator) (err error) {
	if generated, ne := generator.Next(); nil != ne {
		err = ne
	} else {
		t.Id = generated.Get()
	}

	return
}

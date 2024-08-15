package model

import (
	"encoding/json"
	"time"

	"github.com/goexl/model"
)

type Task struct {
	model.Base       `xorm:"extends"`
	model.Optimistic `xorm:"extends"` // 使用乐观锁，优化重试时的性能

	// 下一次重试时间
	// nolint:lll
	Next time.Time `xorm:"datetime notnull index(next) default(CURRENT_TIMESTAMP) comment(一下次更新时间)" json:"next,omitempty"`
	// 重试次数
	Times uint8 `xorm:"tinyint notnull default(0) comment(重试次数)" json:"times,omitempty"`
	// 状态
	// nolint:lll
	Status Status `xorm:"tinyint notnull index(next) default(0) comment(状态，分别是：1、已创建；2、执行中；3、重试中；10、失败；20、成功)" json:"status,omitempty"`
	// 数据
	Data json.RawMessage `xorm:"json null comment(数据)" json:"data,omitempty"`
}

func (*Task) TableComment() string {
	return "任务"
}

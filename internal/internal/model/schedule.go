package model

import (
	"time"

	"github.com/goexl/model"
	"github.com/goexl/task"
)

type Schedule struct {
	model.Base `xorm:"extends"`

	// 目标
	Target uint64 `xorm:"bigint notnull index(target) default(0) comment(目标标识)" json:"target,omitempty"`
	// 类型
	// nolint:lll
	Type task.Type `xorm:"tinyint notnull index(next) default(0) comment(类型，分别是：1、表达式任务；2、固定时间任务；3、周期性任务；4、可被计算的任务；5、只执行一次的任务)" json:"type,omitempty"`
	// 子类型
	// nolint:lll
	Subtype task.Type `xorm:"smallint notnull index(next) default(0) comment(子类型，根据应用自身识别)" json:"subtype,omitempty"`
	// 消耗时间
	Elapsed time.Duration `xorm:"bigint notnull default(0) comment(最大消息时间)"`
	// 重试次数
	Retries uint32 `xorm:"int notnull default(0) comment(重试次数)" json:"retries,omitempty"`
	// 数据
	Data map[string]any `xorm:"json null comment(数据)" json:"data,omitempty"`
}

func (*Schedule) TableComment() string {
	return "计划"
}

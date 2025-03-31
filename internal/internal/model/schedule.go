package model

import (
	"time"

	"github.com/goexl/id"
	"github.com/goexl/model"
	"github.com/goexl/task"
	"github.com/pangum/pangu"
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
	// 超时时间
	Timeout time.Duration `xorm:"bigint notnull default(0) comment(任务超时时间)" json:"timeout,omitempty"`
	// 最大重试次数
	Maximum uint32 `xorm:"int notnull default(0) comment(最大重试次数)" json:"maximum,omitempty"`
	// 数据
	Data map[string]any `xorm:"longtext null comment(数据)" json:"data,omitempty"`
}

func (s *Schedule) BeforeInsert() {
	if 0 == s.Id {
		pangu.New().Get().Dependency().Get(s.setId).Build().Build().Apply()
	}
}

func (*Schedule) TableComment() string {
	return "计划"
}

func (s *Schedule) setId(generator id.Generator) (err error) {
	if generated, ne := generator.Next(); nil != ne {
		err = ne
	} else {
		s.Id = generated.Get()
	}

	return
}

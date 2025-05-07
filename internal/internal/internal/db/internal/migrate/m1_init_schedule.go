package migrate

import (
	"context"

	"github.com/harluo/migrate"
	"github.com/harluo/taskd/internal/internal/internal/model"
	"github.com/harluo/xorm"
)

type M1InitSchedule struct {
	engine *xorm.Engine
}

func newM1InitSchedule(engine *xorm.Engine) migrate.Migration {
	return &M1InitSchedule{
		engine: engine,
	}
}

func (m *M1InitSchedule) Upgrade(_ context.Context) error {
	return m.engine.CreateTables(new(model.Schedule))
}

func (m *M1InitSchedule) Downgrade(_ context.Context) error {
	return m.engine.DropTables(new(model.Schedule))
}

func (*M1InitSchedule) Id() uint64 {
	return 2025_04_30_11_01
}

func (*M1InitSchedule) Description() string {
	return "创建计划表"
}

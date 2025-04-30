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

func (m *M1InitSchedule) Up(_ context.Context) error {
	return m.engine.CreateTables(new(model.Schedule))
}

func (m *M1InitSchedule) Down(_ context.Context) error {
	return m.engine.DropTables(new(model.Schedule))
}

func (*M1InitSchedule) Id() uint64 {
	return 2025_04_30_11_01
}

func (*M1InitSchedule) Module() string {
	return "创建初始化计划表"
}

func (*M1InitSchedule) Description() string {
	return "创建初始化计划表"
}

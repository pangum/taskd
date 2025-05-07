package migrate

import (
	"context"

	"github.com/harluo/migrate"
	"github.com/harluo/taskd/internal/internal/internal/model"
	"github.com/harluo/xorm"
)

type M1InitTask struct {
	engine *xorm.Engine
}

func newM1InitTask(engine *xorm.Engine) migrate.Migration {
	return &M1InitTask{
		engine: engine,
	}
}

func (m *M1InitTask) Upgrade(_ context.Context) error {
	return m.engine.CreateTables(new(model.Task))
}

func (m *M1InitTask) Downgrade(_ context.Context) error {
	return m.engine.DropTables(new(model.Task))
}

func (*M1InitTask) Id() uint64 {
	return 2025_05_07_11_22
}

func (*M1InitTask) Description() string {
	return "创建任务调度细节表"
}

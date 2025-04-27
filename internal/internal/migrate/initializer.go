package migrate

import (
	"context"

	"github.com/harluo/taskd/internal/internal/internal/model"
	"github.com/harluo/xorm"
)

type Initializer struct {
	synchronizer *xorm.Synchronizer
}

func newInitializer(synchronizer *xorm.Synchronizer) *Initializer {
	return &Initializer{
		synchronizer: synchronizer,
	}
}

func (i *Initializer) Initialize(_ context.Context) error {
	return i.synchronizer.Sync(
		new(model.Schedule),
		new(model.Task),
	)
}

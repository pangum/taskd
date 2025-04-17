package model

import (
	"github.com/harluo/xorm"
)

func sync(synchronizer *xorm.Synchronizer) error {
	return synchronizer.Sync(
		new(Schedule),
		new(Task),
	)
}

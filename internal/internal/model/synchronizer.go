package model

import (
	"github.com/pangum/db"
)

type synchronizer struct {
	// 同步器
}

func (*synchronizer) Sync(synchronizer *db.Synchronizer) error {
	return synchronizer.Sync(
		new(Schedule),
		new(Task),
	)
}

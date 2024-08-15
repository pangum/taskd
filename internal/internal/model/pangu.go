package model

import (
	"github.com/pangum/pangu"
)

func init() {
	sync := new(synchronizer)
	pangu.New().Get().Dependency().Get(
		sync.Sync,
	).Build().Build().Apply()
}

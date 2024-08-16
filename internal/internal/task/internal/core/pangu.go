package core

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Put(
		newRunning,
	).Build().Build().Apply()
}

package model

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Get(
		sync,
	).Build().Build().Apply()
}

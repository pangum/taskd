package service

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newTasker,
	).Build().Apply()
}

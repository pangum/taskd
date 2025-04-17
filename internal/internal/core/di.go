package core

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newRunnable,
	).Build().Apply()
}

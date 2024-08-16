package task

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newRunnable,
		newUpstream,
	).Build().Apply()
}

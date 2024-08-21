package schedule

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newRunnable,
	).Build().Apply()
}

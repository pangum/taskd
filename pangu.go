package taskd

import (
	"github.com/pangum/pangu"
	"github.com/pangum/taskd/internal/core"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		internal.NewTasker,
	).Build().Apply()
}

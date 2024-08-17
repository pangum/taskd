package taskd

import (
	"github.com/pangum/pangu"
	"github.com/pangum/taskd/internal"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		internal.NewAgent,
	).Build().Apply()
}

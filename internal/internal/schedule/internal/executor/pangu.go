package executor

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Put(
		newDefault,
	).Build().Build().Apply()
}

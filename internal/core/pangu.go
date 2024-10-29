package core

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newTasker,
	).Build().Apply()
}

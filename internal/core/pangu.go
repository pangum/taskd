package core

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().
		Put(newTasker).Name("pangum.taskd").Group("tasker").Build().
		Build().
		Apply()
}

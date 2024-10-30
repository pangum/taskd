package service

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().
		Put(newTasker).Name("pangum.taskd").Build().
		Build().
		Apply()
}

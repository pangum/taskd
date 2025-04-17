package service

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().
		Put(newTasker).Name("harluo.taskd").Build().
		Build().
		Apply()
}

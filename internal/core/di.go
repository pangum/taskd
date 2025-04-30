package core

import (
	"github.com/harluo/di"
	"github.com/harluo/taskd/internal/internal/put"
	"github.com/harluo/taskd/internal/internal/service"
)

func init() {
	di.New().Instance().Put(
		func(tasker *service.Tasker) put.Tasker {
			return put.Tasker{
				Tasker: tasker,
			}
		},
	).Build().Apply()
}

package service

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newTasker,
	).Build().Apply()
}

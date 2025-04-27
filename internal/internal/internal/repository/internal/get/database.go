package get

import (
	"github.com/goexl/id"
	"github.com/goexl/log"
	"github.com/harluo/di"
	"github.com/harluo/xorm"
)

type Database struct {
	di.Get

	Id     id.Generator
	Logger log.Logger
	DB     *xorm.Engine
}

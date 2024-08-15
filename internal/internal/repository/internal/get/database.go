package get

import (
	"github.com/goexl/id"
	"github.com/goexl/log"
	"github.com/pangum/db"
	"github.com/pangum/pangu"
)

type Database struct {
	pangu.Get

	Id     id.Generator
	Logger log.Logger
	DB     *db.Engine
}

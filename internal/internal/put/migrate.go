package put

import (
	"github.com/harluo/boot"
	"github.com/harluo/di"
)

type Migrate struct {
	di.Put

	Initializer boot.Initializer `group:"initializers"`
}

package get

import (
	"github.com/harluo/xorm"
)

type Tx struct {
	Database

	Tx *xorm.Tx
}

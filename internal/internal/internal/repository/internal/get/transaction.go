package get

import (
	"github.com/harluo/xorm"
)

type Transaction struct {
	Database

	Transaction *xorm.Transaction
}

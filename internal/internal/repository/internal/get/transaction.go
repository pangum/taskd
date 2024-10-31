package get

import (
	"github.com/pangum/db"
)

type Transaction struct {
	Database

	Transaction *db.Transaction
}

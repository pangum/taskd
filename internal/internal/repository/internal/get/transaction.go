package get

import (
	"github.com/pangum/db"
)

type Transaction struct {
	Database

	Tx *db.Transaction
}

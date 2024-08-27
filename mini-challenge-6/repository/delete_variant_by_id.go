package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DELETE_VARIANT_BY_ID_STATEMENT = `DELETE FROM variants WHERE id = $1`

func DeleteVariantById(db *sql.DB, Id int) error {
	_, err := db.Exec(DELETE_VARIANT_BY_ID_STATEMENT, Id)
 
	return err
}

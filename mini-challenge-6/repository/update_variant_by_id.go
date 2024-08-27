package repository

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var UPDATE_VARIANT_BY_ID = `UPDATE variants SET variant_name = $1, quantity = $2, updated_at = $3 WHERE id = $4`

func UpdateVariantById(db *sql.DB, id int, variantName string, quantity int) error {
	_, err := db.Exec(UPDATE_VARIANT_BY_ID, variantName, quantity, time.Now(), id)

	return err
}

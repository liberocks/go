package repository

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var UPDATE_PRODUCT = `UPDATE products SET name = $1, updated_at = $2 WHERE id = $3`

func UpdateProduct(db *sql.DB, id int, name string) error {
	_, err := db.Exec(UPDATE_PRODUCT, name, time.Now(), id)

	return err
}

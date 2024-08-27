package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var CREATE_VARIANT_STATEMENT = `INSERT INTO variants (variant_name, quantity, product_id) VALUES ($1, $2, $3) RETURNING id`

func CreateVariant(db *sql.DB, variantName string, quantity int, productId int) (int, error) {
	var id int
	err := db.QueryRow(CREATE_VARIANT_STATEMENT, variantName, quantity, productId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

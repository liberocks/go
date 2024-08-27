package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var CREATE_PRODUCT_STATEMENT = `INSERT INTO products (name) VALUES ($1) RETURNING id`

func CreateProduct(db *sql.DB, name string) (int, error) {
	var id int
	err := db.QueryRow(CREATE_PRODUCT_STATEMENT, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

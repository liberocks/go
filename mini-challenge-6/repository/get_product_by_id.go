package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/liberocks/go/mini-challenge-6/model"
)

var GET_PRODUCT_BY_ID = `SELECT id, name, created_at, updated_at FROM products WHERE id = $1`

func GetProductById(db *sql.DB, id int) (model.Product, error) {
	var product = model.Product{}

	err := db.QueryRow(GET_PRODUCT_BY_ID, id).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/liberocks/go/mini-challenge-6/model"
)

var GET_PRODUCT_WITH_VARIANT = `SELECT id, variant_name, quantity, product_id, created_at, updated_at FROM variants WHERE product_id = $1`

func GetProductWithVariant(db *sql.DB, id int) (model.Product, error) {
	// Get product by id
	product, err := GetProductById(db, id)
	if err != nil {
		return model.Product{}, err
	}

	// Get variants by product id
	var variants = []model.Variant{}
	rows, err := db.Query(GET_PRODUCT_WITH_VARIANT, id)
	if err != nil {
		return model.Product{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var variant = model.Variant{}
		err = rows.Scan(&variant.ID, &variant.VariantName, &variant.Quantity, &variant.ProductID, &variant.CreatedAt, &variant.UpdatedAt)
		if err != nil {
			return model.Product{}, err
		}
		variants = append(variants, variant)
	}

	product.Variants = variants

	return product, nil
}

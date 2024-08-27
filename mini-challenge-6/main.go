package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/liberocks/go/mini-challenge-6/repository"
	"github.com/liberocks/go/mini-challenge-6/utils"
)

func main() {
	dbUrl := os.Getenv("DB_URL")

	// Run the migration
	utils.Up(dbUrl)

	// Open the database connection
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}

	// Create a product
	productID, err := repository.CreateProduct(db, "Product 1")
	if err != nil {
		panic(err)
	}

	// Get the product by ID
	product, err := repository.GetProductById(db, productID)
	if err != nil {
		panic(err)
	}

	// Print the product
	fmt.Println("Product:")
	fmt.Println(product)

	// Update the product by ID
	err = repository.UpdateProduct(db, productID, "Product 1 Updated")
	if err != nil {
		panic(err)
	}

	// Create one variant
	variantID, err := repository.CreateVariant(db, "Variant 1", 10, productID)
	if err != nil {
		panic(err)
	}

	// Update the variant by ID
	err = repository.UpdateVariantById(db, variantID, "Variant 1 Updated", 20)
	if err != nil {
		panic(err)
	}

	// Create another variant
	_, err = repository.CreateVariant(db, "Variant 2", 16, productID)
	if err != nil {
		panic(err)
	}

	// Get the product with variants
	productWithVariant, err := repository.GetProductWithVariant(db, productID)
	if err != nil {
		panic(err)
	}

	// Print the product with variants
	fmt.Println("Product with Variants:")
	fmt.Println(productWithVariant)

	// Delete the variant by ID
	err = repository.DeleteVariantById(db, variantID)
	if err != nil {
		panic(err)
	}
}

package helpers

import (
	"database/sql"
	"os"
	"sync"
)

var dbOnce sync.Once
var db *sql.DB

func GetDB() *sql.DB {
	// Initialize the database connection
	dbOnce.Do(func() {
		var err error

		dbUrl := os.Getenv("DB_URL")

		// Open the connection
		db, err = sql.Open("postgres", dbUrl)
		if err != nil {
			panic(err)
		}
	})

	return db
}

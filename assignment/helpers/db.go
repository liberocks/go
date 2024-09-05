package helpers

import (
	"database/sql"
	"os"
	"sync"
)

var dbOnce sync.Once
var db *sql.DB

func GetDB() *sql.DB {
	dbOnce.Do(func() {
		var err error

		dbUrl := os.Getenv("DB_URL")

		db, err = sql.Open("postgres", dbUrl)
		if err != nil {
			panic(err)
		}

	})

	return db
}

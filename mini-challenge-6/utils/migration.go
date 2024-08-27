package utils

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Up(dbUrl string) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://sql", "postgres", driver)
	if err != nil {
		panic(err)
	}

	m.Up()
}

func Down(dbUrl string) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://sql", "postgres", driver)
	if err != nil {
		panic(err)
	}

	m.Down()
}

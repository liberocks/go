package utils

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/liberocks/go/assignment/helpers"
)

func Up(dbUrl string) {
	db := helpers.GetDB()

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
	db := helpers.GetDB()

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

package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	// Rpostgresql db
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	//migration
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migratedb() {
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://postgres:0712@localhost:5432/Signup?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

}

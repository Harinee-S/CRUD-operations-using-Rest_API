package handlers

import (
	"fmt"

	//"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"database/sql"
)

const (
	DBUser     = "postgres"
	DBPassword = "0712"
	DBName     = "Signup"
)

// DB set up
func setupDB() *sql.DB {
	var db *sql.DB

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DBUser, DBPassword, DBName)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

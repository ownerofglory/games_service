package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DatabaseUser     = "postgres"
	DatabasePassword = "mypassword"
	DatabaseHost     = "localhost"
	DatabaseName     = "postgres"
)

var (
	db *sql.DB
)

func connectDb() error {
	dbinfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		DatabaseUser,
		DatabasePassword,
		DatabaseHost,
		DatabaseName,
	)

	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		return err
	}

	return db.Ping()
}

package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPosgreSQLDatabaseConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

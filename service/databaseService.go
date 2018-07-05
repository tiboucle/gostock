package service

import (
	"database/sql"
)

const (
	driver   = "sqlite3"
	database = "gostock.db"
)

func InitConnection() (db *sql.DB) {
	db, err := sql.Open(driver, database)
	if err != nil {
		panic(err)
	}
	return db
}

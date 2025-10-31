package database

import (
	"database/sql"
	"log"
)

func InitDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	db.Ping()
	if err != nil {
		return nil, err
	}
	log.Print("Database connected")
	return db, nil
}

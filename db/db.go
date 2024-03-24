package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}

	if err := initStorage(db); err != nil {
		return nil, err
	}

	return db, nil
}

func initStorage(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return err
	}

	return nil
}

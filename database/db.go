package database

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	connStr := "user=postgres password=Curry123! dbname=url_shortener sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("Database is unreachable: %v", err)
	}

	fmt.Println("Successfully connected to the database")
	return nil
}
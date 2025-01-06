package database

import (
	"database/sql"
	"fmt"
	"os"
	_"github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	// Get environment variables for DB connection
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)
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
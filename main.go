package main

import (
	"fmt"
	"log"
	"url-shortener/database"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}

	fmt.Println("Database initialised successfully")
} 

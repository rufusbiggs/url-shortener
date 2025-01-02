package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/database"
	"url-shortener/router"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}
	fmt.Println("Database initialised successfully")

	r := router.InitRouter()
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
} 

package main

import (
	"fmt"
	"log"
	"os"
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback to local 8080 if issues
	}
	log.Printf("Server running on :%s", port)
	log.Fatal(http.ListenAndServe(":" + port, r))
} 

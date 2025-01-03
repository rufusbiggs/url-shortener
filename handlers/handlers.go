package handlers

import (
	"encoding/json"
	"math/rand"
	"fmt"
	"time"
	"net/http"
	"url-shortener/models"
	"url-shortener/database"
	"github.com/gorilla/mux"
)

func generateShortURL() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource((time.Now().UnixNano())))
	b := make([]byte, 6)
	for i := range b {
		b[i] = chars[r.Intn(len(chars))]
	}
	return string(b)
}

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var url models.URL
	_ = json.NewDecoder(r.Body).Decode(&url)

	url.ShortURL = generateShortURL()
	url.CreatedAt = time.Now()

	err := database.DB.QueryRow("INSERT INTO urls (short_url, original_url, created_at) VALUES ($1, $2, $3) RETURNING id", 
		url.ShortURL, url.OriginalURL, url.CreatedAt).Scan(&url.ID)
	if err != nil {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(url)
}

func RedirectToOriginalURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]
	fmt.Println("Recieved short URL: ", shortURL)

	var originalURL string
	var accessCount int
	// Query original URL
	err := database.DB.QueryRow("SELECT original_url, access_count FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL, &accessCount)
	if err != nil {
		fmt.Println("Error fetching from DB:", err)
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	// Increment access count for analytics
	_, err = database.DB.Exec("UPDATE urls SET access_count = $1 WHERE short_url = $2", accessCount+1, shortURL)
	if err != nil {
		fmt.Println("Error updating access count:", err)
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func GetURLAnalytics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	var originalURL string
	var accessCount int
	err := database.DB.QueryRow("SELECT original_url, access_count FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL, &accessCount)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"short_url":      shortURL,
		"original_url":   originalURL,
		"access_count":   accessCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

package main

import (
	"encoding/json"
	"math/rand"
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
	shortURL := vars["shortUrl"]

	var originalURL string
	err := database.DB.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

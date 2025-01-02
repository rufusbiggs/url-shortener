package router

import (
	"github.com/gorilla/mux"
	"url-shortener/handlers"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", handlers.CreateShortURL).Methods("POST")
	router.HandleFunc("/{shortURL}", handlers.RedirectToOriginalURL).Methods("GET")
	return router
}


package main

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", CreateShortURL).Methods("POST")
	router.HandleFunc("/{shortURL}", RedirectToOriginalURL).Methods("GET")
	return router
}


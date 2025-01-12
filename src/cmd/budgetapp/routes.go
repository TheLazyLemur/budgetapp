package main

import (
	"net/http"

	"budgetapp/src/internal/handlers"
)

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HandleIndex)
}

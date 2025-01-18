package main

import (
	"net/http"

	"budgetapp/src/internal/db"
	"budgetapp/src/internal/handlers"
)

func addRoutes(mux *http.ServeMux, qs db.Querier) {
	uHandlers := handlers.NewUserHandlers(qs)

	mux.HandleFunc("GET /", uHandlers.HandleIndex)

	mux.HandleFunc("GET /login", uHandlers.HandleLogin)
	mux.HandleFunc("POST /login", uHandlers.HandleLoginForm)

	mux.HandleFunc("GET /signup", uHandlers.HandleSignup)
	mux.HandleFunc("POST /signup", uHandlers.HandleSignupForm)
}

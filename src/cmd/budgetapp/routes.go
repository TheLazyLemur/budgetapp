package main

import (
	"github.com/go-chi/chi/v5"

	"budgetapp/src/internal/db"
	"budgetapp/src/internal/handlers"
)

func addRoutes(mux *chi.Mux, qs db.Querier) {
	uHandlers := handlers.NewUserHandlers(qs)

	mux.Get("/", uHandlers.HandleIndex)

	mux.Get("/login", uHandlers.HandleLogin)
	mux.Post("/login", uHandlers.HandleLoginForm)

	mux.Get("/signup", uHandlers.HandleSignup)
	mux.Post("/signup", uHandlers.HandleSignupForm)
}

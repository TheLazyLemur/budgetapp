package main

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	"budgetapp/src/internal/db"
	"budgetapp/src/internal/handlers"
)

func addRoutes(mux *chi.Mux, dbc *sql.DB, qs *db.DBTx) {
	uHandlers := handlers.NewUserHandlers(dbc, qs)

	mux.Get("/", uHandlers.HandleIndex)

	mux.Get("/login", uHandlers.HandleLogin)
	mux.Post("/login", uHandlers.HandleLoginForm)

	mux.Get("/signup", uHandlers.HandleSignup)
	mux.Post("/signup", uHandlers.HandleSignupForm)
}

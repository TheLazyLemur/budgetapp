package main

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	"budgetapp/src/internal/db"
	"budgetapp/src/internal/handlers"
	"budgetapp/src/internal/middleware"
)

func addRoutes(mux *chi.Mux, dbc *sql.DB, qs *db.DBTx) {
	uHandlers := handlers.NewUserHandlers(dbc, qs)

	mux.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(qs))
		r.Get("/", uHandlers.HandleIndex)
	})

	mux.Get("/login", uHandlers.HandleLogin)
	mux.Post("/login", uHandlers.HandleLoginForm)

	mux.Get("/signup", uHandlers.HandleSignup)
	mux.Post("/signup", uHandlers.HandleSignupForm)
}

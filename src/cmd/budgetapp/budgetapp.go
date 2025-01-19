package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"

	"budgetapp/src/internal/db"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbc, err := sql.Open("sqlite3", "budgetapp.db")
	if err != nil {
		panic(err)
	}
	defer dbc.Close()

	if err := dbc.Ping(); err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	addRoutes(r, db.New(dbc))

	slog.Info("starting server", slog.String("address", port))

	if err := http.ListenAndServe(":"+port, r); err != nil {
		slog.Error("something went wrong", slog.String("err", err.Error()))
	}
}

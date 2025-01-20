package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"budgetapp/src/internal/db"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	url := os.Getenv("DB_URL") + os.Getenv("AUTH_TOKEN")
	dbc, err := sql.Open("libsql", url)
	if err != nil {
		panic(err)
	}
	defer dbc.Close()

	if err := dbc.Ping(); err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	qs := db.NewDB(dbc)

	addRoutes(r, dbc, qs)

	slog.Info("starting server", slog.String("address", port))

	if err := http.ListenAndServe(":"+port, r); err != nil {
		slog.Error("something went wrong", slog.String("err", err.Error()))
	}
}

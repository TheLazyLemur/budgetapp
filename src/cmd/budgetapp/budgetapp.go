package main

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"

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
	qs := db.New(dbc)
	user, err := qs.InsertUser(context.Background(), db.InsertUserParams{
		Name:           "John Doe",
		Email:          "john@doe.com",
		HashedPassword: "123456",
	})
	if err != nil {
		slog.Info("Default user already exists")
	} else {
		slog.Info("Default user created", slog.String("user", user.UserID))
	}

	mux := http.NewServeMux()

	addRoutes(mux)

	slog.Info("starting server", slog.String("address", port))

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		slog.Error("something went wrong", slog.String("err", err.Error()))
	}
}

package main

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"budgetapp"
)

func main() {
	godotenv.Load()

	url := os.Getenv("DB_URL") + os.Getenv("AUTH_TOKEN")
	dbc, err := sql.Open("libsql", url)
	if err != nil {
		panic(err)
	}
	defer dbc.Close()

	goose.SetBaseFS(budgetapp.EmbedMigrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}

	if err := goose.Up(dbc, "db/migrations"); err != nil {
		panic(err)
	}
}

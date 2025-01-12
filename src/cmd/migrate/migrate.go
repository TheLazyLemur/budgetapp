package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"

	"budgetapp"
)

func main() {
	dbc, err := sql.Open("sqlite3", "file:budgetapp.db?cache=shared&mode=rwc")
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

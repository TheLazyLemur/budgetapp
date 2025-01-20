package main

import (
	"context"
	"database/sql"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"budgetapp/src/internal/db"
)

func main() {
	godotenv.Load()

	url := os.Getenv("DB_URL") + os.Getenv("AUTH_TOKEN")
	dbc, err := sql.Open("libsql", url)
	if err != nil {
		panic(err)
	}
	defer dbc.Close()

	qs := db.NewDB(dbc)

	err = qs.CreateTransaction(context.Background(), db.CreateTransactionParams{
		ID:     uuid.Must(uuid.NewRandom()).String(),
		UserID: "d48255b4-5cd6-4594-ade0-b794b8dfaf1d",
		Amount: 100,
		Type:   "income",
		Description: sql.NullString{
			String: "sdasd",
			Valid:  true,
		},
	})
	if err != nil {
		panic(err)
	}
}

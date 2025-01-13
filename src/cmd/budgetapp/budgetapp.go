package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	addRoutes(mux)

	slog.Info("starting server", slog.String("address", port))

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		slog.Error("something went wrong", slog.String("err", err.Error()))
	}
}

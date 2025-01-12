package main

import (
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	addRoutes(mux)

	slog.Info("starting server", slog.String("address", ":8080"))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("something went wrong", slog.String("err", err.Error()))
	}
}

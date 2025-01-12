package handlers

import (
	"log/slog"
	"net/http"

	"budgetapp/src/internal/views"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if err := views.Index().Render(r.Context(), w); err != nil {
		slog.Error("failed to render", slog.String("err", err.Error()))
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return
	}
}

package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"budgetapp/src/internal/core"
	"budgetapp/src/internal/db"
)

func AuthMiddleware(queries *db.DBTx) func(next http.Handler) http.Handler {
	return authMiddleware(queries)
}

func authMiddleware(queries *db.DBTx) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			if err != nil {
				slog.Error("failed to get session cookie", slog.String("err", err.Error()))
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			user, err := core.GetUserBySessionID(r.Context(), queries, cookie.Value)
			if err != nil {
				slog.Error("failed to get user", slog.String("err", err.Error()))
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

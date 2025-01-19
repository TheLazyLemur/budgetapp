package handlers

import (
	"database/sql"
	"log/slog"
	"net/http"

	"budgetapp/src/internal/core"
	"budgetapp/src/internal/db"
	"budgetapp/src/internal/views"
)

type UserHandlers struct {
	queries *db.DBTx
}

func NewUserHandlers(dbc *sql.DB, queries *db.DBTx) *UserHandlers {
	return &UserHandlers{
		queries: queries,
	}
}

func (h *UserHandlers) HandleIndex(w http.ResponseWriter, r *http.Request) {
	if err := views.IndexPage().Render(r.Context(), w); err != nil {
		slog.Error("failed to render", slog.String("err", err.Error()))
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandlers) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if err := views.LoginPage(nil).Render(r.Context(), w); err != nil {
		slog.Error("failed to render", slog.String("err", err.Error()))
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandlers) HandleLoginForm(w http.ResponseWriter, r *http.Request) {
	loginRequest, err := loginRequestFromForm(r)
	if err != nil {
		errors := []string{err.Error()}
		if err := views.LoginPage(errors).Render(r.Context(), w); err != nil {
			slog.Error("failed to render", slog.String("err", err.Error()))
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
		return
	}

	_, err = core.LoginUser(r.Context(), h.queries, loginRequest.Email, loginRequest.Password)
	if err != nil {
		errors := []string{err.Error()}
		if err := views.LoginPage(errors).Render(r.Context(), w); err != nil {
			slog.Error("failed to render", slog.String("err", err.Error()))
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
		return
	}

	// TODO: Set cookie with session ID and expiration date

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *UserHandlers) HandleSignup(w http.ResponseWriter, r *http.Request) {
	if err := views.SignupPage(nil).Render(r.Context(), w); err != nil {
		slog.Error("failed to render", slog.String("err", err.Error()))
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandlers) HandleSignupForm(w http.ResponseWriter, r *http.Request) {
	signupRequest, err := signupRequestFromForm(r)
	if err != nil {
		errors := []string{err.Error()}
		if err := renderSignupPage(w, r, errors); err != nil {
			return
		}
		return
	}

	_, err = core.RegisterUser(
		r.Context(),
		h.queries,
		signupRequest.Name,
		signupRequest.Email,
		signupRequest.Password,
	)
	if err != nil {
		errors := []string{err.Error()}
		if err := renderSignupPage(w, r, errors); err != nil {
			return
		}
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func renderSignupPage(w http.ResponseWriter, r *http.Request, errors []string) error {
	if err := views.SignupPage(errors).Render(r.Context(), w); err != nil {
		slog.Error("failed to render", slog.String("err", err.Error()))
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return err
	}

	return nil
}

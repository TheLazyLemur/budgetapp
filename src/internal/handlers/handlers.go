package handlers

import (
	"database/sql"
	"log/slog"
	"net/http"
	"time"

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
	user, ok := r.Context().Value("user").(db.User)
	if !ok {
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}

	if err := views.IndexPage(user).Render(r.Context(), w); err != nil {
		slog.Error("failed to render", slog.String("err", err.Error()))
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandlers) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if err := views.LoginPage(db.User{}, nil).Render(r.Context(), w); err != nil {
		slog.Error("failed to render", slog.String("err", err.Error()))
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandlers) HandleLoginForm(w http.ResponseWriter, r *http.Request) {
	loginRequest, err := loginRequestFromForm(r)
	if err != nil {
		errors := []string{err.Error()}
		if err := views.LoginPage(db.User{}, errors).Render(r.Context(), w); err != nil {
			slog.Error("failed to render", slog.String("err", err.Error()))
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
		return
	}

	loginResult, err := core.LoginUser(
		r.Context(),
		h.queries,
		loginRequest.Email,
		loginRequest.Password,
	)
	if err != nil {
		errors := []string{err.Error()}
		if err := views.LoginPage(db.User{}, errors).Render(r.Context(), w); err != nil {
			slog.Error("failed to render", slog.String("err", err.Error()))
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    loginResult.Token,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,
		Secure:   true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *UserHandlers) HandleSignup(w http.ResponseWriter, r *http.Request) {
	if err := views.SignupPage(db.User{}, nil).Render(r.Context(), w); err != nil {
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

	registerResult, err := core.RegisterUser(
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

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    registerResult.Token,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,
		Secure:   true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func renderSignupPage(w http.ResponseWriter, r *http.Request, errors []string) error {
	if err := views.SignupPage(db.User{}, errors).Render(r.Context(), w); err != nil {
		slog.Error("failed to render", slog.String("err", err.Error()))
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return err
	}

	return nil
}

func (h *UserHandlers) HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24 * 7),
		HttpOnly: true,
		Secure:   true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

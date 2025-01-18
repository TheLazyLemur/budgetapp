package handlers

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"budgetapp/src/internal/db"
	"budgetapp/src/internal/views"
)

type UserHandlers struct {
	queries db.Querier
}

func NewUserHandlers(queries db.Querier) *UserHandlers {
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

	usr, err := h.queries.GetUserByEmailAndHashedPassword(
		r.Context(),
		db.GetUserByEmailAndHashedPasswordParams{
			Email:          loginRequest.Email,
			HashedPassword: loginRequest.Password,
		},
	)
	if err != nil {
		errors := []string{err.Error()}
		if err := views.LoginPage(errors).Render(r.Context(), w); err != nil {
			slog.Error("failed to render", slog.String("err", err.Error()))
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
		return
	}

	_ = usr
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

	_, err = h.queries.InsertUser(r.Context(), db.InsertUserParams{
		UserID:         uuid.Must(uuid.NewRandom()).String(),
		Name:           signupRequest.Name,
		Email:          signupRequest.Email,
		HashedPassword: signupRequest.Password,
	})
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

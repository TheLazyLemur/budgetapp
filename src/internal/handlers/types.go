package handlers

import (
	"errors"
	"log/slog"
	"net/http"
)

type signupRequest struct {
	Name     string
	Email    string
	Password string
}

func signupRequestFromForm(r *http.Request) (signupRequest, error) {
	name := r.FormValue("name")
	if len(name) == 0 {
		return signupRequest{}, errors.New("name is required")
	}

	email := r.FormValue("email")
	if len(email) == 0 {
		return signupRequest{}, errors.New("email is required")
	}

	password := r.FormValue("password")
	if len(password) == 0 {
		return signupRequest{}, errors.New("password is required")
	}

	slog.Info(
		"Sign up form",
		slog.String("name", name),
		slog.String("email", email),
		slog.String("password", password),
	)

	return signupRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}

type loginRequest struct {
	Email    string
	Password string
}

func loginRequestFromForm(r *http.Request) (loginRequest, error) {
	email := r.FormValue("email")
	if len(email) == 0 {
		return loginRequest{}, errors.New("email is required")
	}

	password := r.FormValue("password")
	if len(password) == 0 {
		return loginRequest{}, errors.New("password is required")
	}

	slog.Info(
		"Login form",
		slog.String("email", email),
		slog.String("password", password),
	)

	return loginRequest{
		Email:    email,
		Password: password,
	}, nil
}

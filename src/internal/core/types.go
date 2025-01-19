package core

import (
	"errors"

	"budgetapp/src/internal/db"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrSomethingWentWrong = errors.New("something went wrong")
)

type UserLoginResult struct {
	User  db.User
	Token string
}

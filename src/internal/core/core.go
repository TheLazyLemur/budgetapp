package core

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"budgetapp/src/internal/db"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func comparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func LoginUser(
	ctx context.Context,
	repo *db.DBTx,
	email, password string,
) (result UserLoginResult, err error) {
	repoWithTx, err := repo.BeginTx(ctx)
	if err != nil {
		return UserLoginResult{}, err
	}
	defer func() {
		if err != nil {
			repoWithTx.Rollback()
		} else {
			err = repoWithTx.Commit()
		}
	}()

	usr, err := repoWithTx.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return UserLoginResult{}, ErrInvalidCredentials
		}
		slog.Error("failed to get user", slog.String("err", err.Error()))
		return UserLoginResult{}, ErrSomethingWentWrong
	}

	if !comparePassword(password, usr.HashedPassword) {
		return UserLoginResult{}, ErrInvalidCredentials
	}

	sessionID := uuid.Must(uuid.NewRandom()).String()
	err = repoWithTx.CreateSession(ctx, db.CreateSessionParams{
		SessionID: sessionID,
		UserID:    usr.UserID,
	})
	if err != nil {
		slog.Error("failed to create session", slog.String("err", err.Error()))
		return UserLoginResult{}, ErrSomethingWentWrong
	}

	return UserLoginResult{
		User:  usr,
		Token: sessionID,
	}, nil
}

func RegisterUser(
	ctx context.Context,
	repo *db.DBTx,
	name, email, password string,
) (result UserRegisterResult, err error) {
	repoWithTx, err := repo.BeginTx(ctx)
	if err != nil {
		return UserRegisterResult{}, err
	}
	defer func() {
		if err != nil {
			repoWithTx.Rollback()
		} else {
			err = repoWithTx.Commit()
		}
	}()

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return UserRegisterResult{}, err
	}

	usr, err := repoWithTx.InsertUser(ctx, db.InsertUserParams{
		UserID:         uuid.Must(uuid.NewRandom()).String(),
		Name:           name,
		Email:          email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		slog.Error("failed to insert user", slog.String("err", err.Error()))
		return UserRegisterResult{}, ErrSomethingWentWrong
	}

	sessionID := uuid.Must(uuid.NewRandom()).String()
	err = repoWithTx.CreateSession(ctx, db.CreateSessionParams{
		SessionID: sessionID,
		UserID:    usr.UserID,
	})
	if err != nil {
		slog.Error("failed to create session", slog.String("err", err.Error()))
		return UserRegisterResult{}, ErrSomethingWentWrong
	}

	return UserRegisterResult{
		User:  usr,
		Token: sessionID,
	}, nil
}

func GetUserBySessionID(
	ctx context.Context,
	repo *db.DBTx,
	sessionID string,
) (db.User, error) {
	session, err := repo.GetSessionByID(ctx, sessionID)
	if err != nil {
		return db.User{}, err
	}

	user, err := repo.GetUserByID(ctx, session.UserID)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

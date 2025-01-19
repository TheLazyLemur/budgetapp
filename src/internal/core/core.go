package core

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/google/uuid"

	"budgetapp/src/internal/db"
)

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

	usr, err := repoWithTx.GetUserByEmailAndHashedPassword(
		ctx,
		db.GetUserByEmailAndHashedPasswordParams{
			Email:          email,
			HashedPassword: password,
		},
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return UserLoginResult{}, ErrInvalidCredentials
		}
		slog.Error("failed to get user", slog.String("err", err.Error()))
		return UserLoginResult{}, ErrSomethingWentWrong
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

	usr, err := repoWithTx.InsertUser(ctx, db.InsertUserParams{
		UserID:         uuid.Must(uuid.NewRandom()).String(),
		Name:           name,
		Email:          email,
		HashedPassword: password,
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

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	// Delete a user by ID and hashed password
	DeleteUserByID(ctx context.Context, arg DeleteUserByIDParams) error
	// Get a user by email and hashed password
	GetUserByEmailAndHashedPassword(ctx context.Context, arg GetUserByEmailAndHashedPasswordParams) (User, error)
	// Inset a new user
	InsertUser(ctx context.Context, arg InsertUserParams) (User, error)
	// Update a user by ID, this will also update the updated_at field in the database
	UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) error
}

var _ Querier = (*Queries)(nil)

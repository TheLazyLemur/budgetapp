// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
)

const createSession = `-- name: CreateSession :exec
INSERT INTO sessions (session_id, user_id, expires) 
VALUES (?, ?, datetime('now', '+1 hour'))
`

type CreateSessionParams struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
}

// Create a new session for a user
func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) error {
	_, err := q.db.ExecContext(ctx, createSession, arg.SessionID, arg.UserID)
	return err
}

const deleteSessionByID = `-- name: DeleteSessionByID :exec
DELETE FROM sessions 
WHERE session_id = ?
`

// Delete a session by ID
func (q *Queries) DeleteSessionByID(ctx context.Context, sessionID string) error {
	_, err := q.db.ExecContext(ctx, deleteSessionByID, sessionID)
	return err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE FROM users 
WHERE user_id = ? 
AND hashed_password = ?
`

type DeleteUserByIDParams struct {
	UserID         string `json:"user_id"`
	HashedPassword string `json:"hashed_password"`
}

// Delete a user by ID and hashed password
func (q *Queries) DeleteUserByID(ctx context.Context, arg DeleteUserByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserByID, arg.UserID, arg.HashedPassword)
	return err
}

const getSessionByID = `-- name: GetSessionByID :one
SELECT session_id, user_id, expires, date_created FROM sessions 
WHERE session_id = ? 
LIMIT 1
`

// Get a session by ID
func (q *Queries) GetSessionByID(ctx context.Context, sessionID string) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByID, sessionID)
	var i Session
	err := row.Scan(
		&i.SessionID,
		&i.UserID,
		&i.Expires,
		&i.DateCreated,
	)
	return i, err
}

const getUserByEmailAndHashedPassword = `-- name: GetUserByEmailAndHashedPassword :one
SELECT user_id, name, email, hashed_password, date_created, date_updated FROM users 
WHERE email = ? 
AND hashed_password = ? 
LIMIT 1
`

type GetUserByEmailAndHashedPasswordParams struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

// Get a user by email and hashed password
func (q *Queries) GetUserByEmailAndHashedPassword(ctx context.Context, arg GetUserByEmailAndHashedPasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmailAndHashedPassword, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.DateCreated,
		&i.DateUpdated,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users (user_id, name, email, hashed_password) 
VALUES (?, ?, ?, ?)
RETURNING user_id, name, email, hashed_password, date_created, date_updated
`

type InsertUserParams struct {
	UserID         string `json:"user_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

// Inset a new user
func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.UserID,
		arg.Name,
		arg.Email,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.DateCreated,
		&i.DateUpdated,
	)
	return i, err
}

const updateUserByID = `-- name: UpdateUserByID :exec
UPDATE users 
SET name = ?, 
    email = ?, 
    hashed_password = ?, 
    updated_at = datetime('now') 
WHERE user_id = ?
`

type UpdateUserByIDParams struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	UserID         string `json:"user_id"`
}

// Update a user by ID, this will also update the updated_at field in the database
func (q *Queries) UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) error {
	_, err := q.db.ExecContext(ctx, updateUserByID,
		arg.Name,
		arg.Email,
		arg.HashedPassword,
		arg.UserID,
	)
	return err
}
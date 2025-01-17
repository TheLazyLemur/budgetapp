-- Inset a new user
-- name: InsertUser :one
INSERT INTO users (user_id, name, email, hashed_password) 
VALUES (?, ?, ?, ?)
RETURNING *;

-- Get a user by email and hashed password
-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = ? 
LIMIT 1;

-- Get a user by ID
-- name: GetUserByID :one
SELECT * FROM users 
WHERE user_id = ? 
LIMIT 1;

-- Update a user by ID, this will also update the updated_at field in the database
-- name: UpdateUserByID :exec
UPDATE users 
SET name = ?, 
    email = ?, 
    hashed_password = ?, 
    updated_at = datetime('now') 
WHERE user_id = ?;

-- Delete a user by ID and hashed password
-- name: DeleteUserByID :exec
DELETE FROM users 
WHERE user_id = ? 
AND hashed_password = ?;

-- Create a new session for a user
-- name: CreateSession :exec
INSERT INTO sessions (session_id, user_id, expires) 
VALUES (?, ?, datetime('now', '+1 hour'));

-- Get a session by ID
-- name: GetSessionByID :one
SELECT * FROM sessions 
WHERE session_id = ? 
LIMIT 1;

-- Delete a session by ID
-- name: DeleteSessionByID :exec
DELETE FROM sessions 
WHERE session_id = ?;

-- Create a new transaction for a user
-- name: CreateTransaction :exec
INSERT INTO transactions (id, user_id, amount, type, description, created_at)
VALUES (?, ?, ?, ?, ?, datetime('now'));

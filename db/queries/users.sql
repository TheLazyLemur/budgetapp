-- Inset a new user
-- name: InsertUser :one
INSERT INTO users (user_id, name, email, hashed_password) 
VALUES (?, ?, ?, ?)
RETURNING *;

-- Get a user by email and hashed password
-- name: GetUserByEmailAndHashedPassword :one
SELECT * FROM users 
WHERE email = ? 
AND hashed_password = ? 
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

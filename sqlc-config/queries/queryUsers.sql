-- name: CreateUser :one
INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, username;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id DESC;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT id, username, password 
FROM users 
WHERE username = $1;

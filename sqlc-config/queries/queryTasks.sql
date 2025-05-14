-- name: CreateTask :one
INSERT INTO tasks (text, created_at)
VALUES ($1, $2)
RETURNING *;

-- name: GetTaskByID :one
SELECT * 
FROM tasks 
WHERE id = $1
ORDER BY created_at;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

-- name: UpdateTask :exec
UPDATE tasks
SET text = $2
WHERE id = $1;

-- name: GetAllTasks :many
SELECT * FROM tasks;

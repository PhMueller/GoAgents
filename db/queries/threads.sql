-- name: CreateThread :one
INSERT INTO threads (id, title) VALUES ($1, $2)
RETURNING *;

-- name: GetThread :one
SELECT * FROM threads WHERE id = $1;

-- name: GetThreads :many
SELECT * FROM threads WHERE id in ($1) ORDER BY id;

-- name: DeleteThread :one
DELETE FROM threads WHERE id = $1 RETURNING *;
-- name: CreateThread :one
INSERT INTO threads (title) VALUES ( $1) RETURNING *;

-- name: GetThreadById :one
SELECT * FROM threads WHERE id = $1;

-- name: DeleteThread :one
DELETE FROM threads WHERE id = $1 RETURNING *;

-- name: ListThreads :many
SELECT * FROM threads t order by t.pk;
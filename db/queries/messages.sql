-- name: GetMessage :one
SELECT * FROM messages WHERE id = $1;

-- name: GetMessages :many
SELECT * FROM messages WHERE id in ($1) ORDER BY id;

-- name: GetMessageByMessageId :one
SELECT * FROM messages WHERE id = $1;

-- name: GetMessagesByThreadId :many
SELECT * FROM messages WHERE thread_id = $1 ORDER BY id;

-- name: CreateMessage :one
INSERT INTO messages (thread_id, content) VALUES ($1, $2) RETURNING *;

-- name: DeleteMessage :one
DELETE FROM messages WHERE id = $1 RETURNING *;

-- name: CreateEvent :one
INSERT INTO events (name, created_at, updated_at, event_date)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetEvent :one
SELECT * FROM events WHERE id = $1;

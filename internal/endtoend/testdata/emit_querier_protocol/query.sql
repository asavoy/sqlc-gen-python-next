-- name: GetAuthor :one
SELECT id, name, bio FROM authors WHERE id = $1;

-- name: ListAuthors :many
SELECT id, name, bio FROM authors ORDER BY name;

-- name: CreateAuthor :exec
INSERT INTO authors (name, bio) VALUES ($1, $2);

-- name: DeleteAuthor :execrows
DELETE FROM authors WHERE id = $1;

-- name: GetArticle :one
SELECT * FROM articles WHERE id = $1;

-- name: CreateArticle :one
INSERT INTO articles (metadata, settings, author_id)
VALUES ($1, $2, $3)
RETURNING *;

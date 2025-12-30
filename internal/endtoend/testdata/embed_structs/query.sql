-- name: GetBookWithAuthor :one
SELECT
    sqlc.embed(books),
    sqlc.embed(authors)
FROM books
JOIN authors ON books.author_id = authors.id
WHERE books.id = $1;

-- name: ListBooksWithAuthors :many
SELECT
    sqlc.embed(books),
    sqlc.embed(authors)
FROM books
JOIN authors ON books.author_id = authors.id
ORDER BY books.title;

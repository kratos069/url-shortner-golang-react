-- name: CreateURL :one
INSERT INTO urls (code, original_url)
VALUES ($1, $2)
RETURNING *;

-- name: GetURLByCode :one
SELECT * FROM urls
WHERE code = $1;

-- name: CheckCodeExists :one
SELECT 1 FROM urls
WHERE code = $1;
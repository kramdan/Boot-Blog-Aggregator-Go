-- name: GetUserByID :one
SELECT users.name FROM users WHERE users.id = $1;
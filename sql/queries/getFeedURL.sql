-- name: GetFeedURL :one
SELECT * FROM feeds WHERE feeds.Url = $1;
-- name: NextFeedToFetch :one
SELECT * FROM feeds ORDER BY last_fetched_at NULLS FIRST;
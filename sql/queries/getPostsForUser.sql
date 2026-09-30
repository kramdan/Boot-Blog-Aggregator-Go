-- name: GetPostsForUser :many
SELECT * FROM posts
INNER JOIN feed_follows
    ON posts.feed_id = feed_follows.feed_id
WHERE posts.feed_id = feed_follows.feed_id
ORDER BY published_at DESC
LIMIT $1;

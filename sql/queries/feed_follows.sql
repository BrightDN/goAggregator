-- name: CreateFeedFollow :one

WITH inserted AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT
  i.id,
  i.created_at,
  i.updated_at,
  u.name AS username,
  f.name AS feedname
FROM inserted i
JOIN users u ON u.id = i.user_id
JOIN feeds f ON f.id = i.feed_id;


-- name: GetFeedFollowsForUser :many
SELECT 
    feed_follows.id,
    feed_follows.created_at,
    feed_follows.updated_at,
    users.name AS user,
    feeds.name AS feedname
FROM feed_follows
JOIN users ON users.id = feed_follows.user_id
JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;

-- name: DeleteByUrl :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;

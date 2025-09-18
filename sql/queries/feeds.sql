-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES(
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT
Name,
Url,
(SELECT users.Name
	FROM users
	WHERE users.id = feeds.user_id) AS Username
FROM feeds;

-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1;

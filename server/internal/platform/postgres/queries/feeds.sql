-- name: CheckFeedExists :one
SELECT EXISTS (
    SELECT 1 FROM feeds
    WHERE id = $1
);

-- name: GetFeeds :many
SELECT 
    f.id,
    f.title,
    f.content,
    f.feed_image,
    u.fullname,
    u.username,
    u.avatar_url,
    (SELECT COUNT(*) FROM likes l WHERE l.feed_id = f.id) AS like_count,
    (SELECT COUNT(*) FROM comments c WHERE c.feed_id = f.id) AS comment_count,
    EXISTS (
        SELECT 1 FROM likes l 
        WHERE l.feed_id = f.id 
        AND l.user_id = sqlc.narg('viewer_id')::uuid
        AND l.entity_type = 'feed'
    ) AS is_liked,
    f.created_at
FROM feeds f
LEFT JOIN users u ON u.id = f.user_id
WHERE
    (
        (sqlc.narg('cursor_date')::timestamptz IS NULL 
         OR sqlc.narg('cursor_id')::uuid IS NULL)
        OR (
            f.created_at < sqlc.narg('cursor_date')::timestamptz
            OR (
                f.created_at = sqlc.narg('cursor_date')::timestamptz 
                AND f.id < sqlc.narg('cursor_id')::uuid
            )
        )
    )
ORDER BY f.created_at DESC, f.id DESC
LIMIT sqlc.arg('limit');


-- name: FeedDetails :one
SELECT 
  f.id,
  f.created_at,
  f.title,
  f.content,
  f.feed_image,
  u.avatar_url,
  u.fullname,
  u.username,
  (SELECT COUNT(*) FROM likes l WHERE l.feed_id = f.id) AS like_count,
  (SELECT COUNT(*) FROM comments c WHERE c.feed_id = f.id) AS comment_count,
  EXISTS (
      SELECT 1 FROM likes l 
      WHERE l.feed_id = f.id 
      AND l.user_id = sqlc.narg('viewer_id')
      AND l.entity_type = 'feed'
  ) AS is_liked
FROM feeds f
LEFT JOIN users u ON u.id = f.user_id
WHERE f.id = sqlc.arg('id');


-- name: CreateFeed :exec
INSERT INTO feeds (
    user_id,
    title,
    content,
    feed_image
) 
VALUES ($1, $2, $3, $4);


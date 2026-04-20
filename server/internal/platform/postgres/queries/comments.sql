-- name: FeedComments :many
SELECT
    c.id,
    c.parent_id,
    c.feed_id,
    c.created_at,
    c.content,
    u.avatar_url,
    u.fullname,
    u.username,
    (SELECT COUNT(*) FROM likes l WHERE l.comment_id = c.id) AS like_count,
    EXISTS (
        SELECT 1 FROM likes l 
        WHERE l.comment_id = c.id
        AND l.user_id = sqlc.narg('viewer_id')::uuid
        AND entity_type = 'comment'
    ) AS is_liked
FROM comments c
LEFT JOIN users u ON u.id = c.author_id
WHERE c.feed_id = @feed_id
AND (
    sqlc.narg('cursor_date')::TIMESTAMPTZ IS NULL
    OR sqlc.narg('cursor_id')::UUID IS NULL
    OR c.created_at < sqlc.narg('cursor_date')::TIMESTAMPTZ
    OR (c.created_at = sqlc.narg('cursor_date')::TIMESTAMPTZ AND c.id < sqlc.narg('cursor_id')::UUID)
)
ORDER BY c.created_at DESC, c.id DESC
LIMIT @limit_count;


-- name: CheckCommentExists :one
SELECT EXISTS (
    SELECT 1 FROM comments
    WHERE id = @id
);


-- name: CreateComment :exec
INSERT INTO comments (
    content,
    author_id,
    feed_id,
    parent_id
)
VALUES (@content, @author_id, @feed_id, @parent_id);


-- name: GetFeedCommentByID :one
SELECT
    c.id,
    c.parent_id,
    c.feed_id,
    c.content,
    c.created_at,
    u.fullname,
    u.username,
    u.avatar_url,
    (SELECT COUNT(*) FROM likes l WHERE l.comment_id = c.id) AS like_count,
    EXISTS (
        SELECT 1 FROM likes l 
        WHERE l.comment_id = c.id 
        AND l.user_id = sqlc.narg('viewer_id')::uuid
    ) AS is_liked
FROM comments c
LEFT JOIN users u ON u.id = c.author_id
WHERE c.id = @id;


-- name: GetCommentByID :one
SELECT *
FROM comments
WHERE id = @id;
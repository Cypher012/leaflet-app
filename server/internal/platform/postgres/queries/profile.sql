-- name: GetUserProfileByUsername :one
SELECT
    u.id,
    u.fullname,
    u.username,
    u.bio,
    u.avatar_url,
    (SELECT COUNT(*) FROM likes l WHERE l.user_id = u.id) AS like_count,
    (SELECT COUNT(*) FROM comments c WHERE c.author_id = u.id) AS comment_count,
    (SELECT COUNT(*) FROM feeds ca WHERE ca.user_id = u.id) AS feed_count
FROM users u
WHERE u.username = $1;

-- name: UpdateUserProfile :one
UPDATE users
SET
    fullname = $1,
    username = $2,
    bio = $3
WHERE id = $4
RETURNING *;


-- name: UserActivity :many
WITH activity AS (
    SELECT
        'feed'        AS type,
        f.id,
        f.title       AS title,
        f.content     AS content,
        f.feed_image  AS feed_image,
        ''::TEXT    AS comment_body,
        ''::TEXT    AS parent_feed_title,
        (SELECT COUNT(*) FROM likes l WHERE l.feed_id = f.id)       AS like_count,
        (SELECT COUNT(*) FROM comments c WHERE c.feed_id = f.id)  AS comment_count,
        EXISTS (
            SELECT 1 FROM likes l
            WHERE l.feed_id = f.id
            AND  l.user_id = sqlc.narg('viewer_id')::uuid
            AND entity_type = 'feed'
        ) AS is_liked,
        f.created_at
    FROM feeds f
    WHERE f.user_id = @user_id

    UNION ALL

    SELECT
        'comment'     AS type,
        c.id,
        ''::TEXT    AS title,
        ''::TEXT    AS content,
        ''::TEXT    AS feed_image,
        c.content    AS comment_body,
        f.title       AS parent_feed_title,
        (SELECT COUNT(*) FROM likes l WHERE l.comment_id = c.id)   AS like_count,
        0::BIGINT  AS comment_count,
        EXISTS (
            SELECT 1 FROM likes l
            WHERE l.comment_id = c.id
            AND l.user_id = sqlc.narg('viewer_id')::uuid
            AND entity_type = 'comment'
        ) AS is_liked,
        c.created_at
    FROM comments c
    JOIN feeds f ON c.feed_id = f.id
    WHERE c.author_id = @user_id
)

SELECT *
FROM activity
WHERE (
    sqlc.narg('cursor_date')::TIMESTAMPTZ IS NULL
    OR sqlc.narg('cursor_id')::UUID IS NULL
    OR created_at < sqlc.narg('cursor_date')::TIMESTAMPTZ
    OR (created_at = sqlc.narg('cursor_date')::TIMESTAMPTZ AND id < sqlc.narg('cursor_id')::UUID)
)
ORDER BY created_at DESC, id DESC
LIMIT sqlc.arg('limit');

-- name: UserComments :many
SELECT
    c.id,
    f.title,
    c.content,
    c.created_at,
    (SELECT COUNT(*) FROM likes l WHERE l.comment_id = c.id) AS like_count,
    EXISTS (
        SELECT 1 FROM likes l
        WHERE l.comment_id = c.id
        AND l.user_id = sqlc.narg('viewer_id')::uuid
    ) AS is_liked
FROM comments c
LEFT JOIN feeds f ON f.id = c.feed_id
WHERE c.author_id = @user_id
AND (
    sqlc.narg('cursor_date')::TIMESTAMPTZ IS NULL
    OR sqlc.narg('cursor_id')::UUID IS NULL
    OR c.created_at < sqlc.narg('cursor_date')::TIMESTAMPTZ
    OR (
        c.created_at = sqlc.narg('cursor_date')::TIMESTAMPTZ
        AND c.id < sqlc.narg('cursor_id')::UUID
    )
)
ORDER BY c.created_at DESC, c.id DESC
LIMIT sqlc.arg('limit');


-- name: UserFeeds :many
SELECT
    f.id,
    f.title,
    f.content,
    f.feed_image,
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
WHERE
    f.user_id = sqlc.arg('user_id')::uuid
    AND (
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

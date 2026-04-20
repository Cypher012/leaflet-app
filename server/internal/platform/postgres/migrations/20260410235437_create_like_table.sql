-- +goose Up
CREATE TABLE IF NOT EXISTS likes (
    id          UUID        PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id     UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id     UUID        REFERENCES feeds(id) ON DELETE CASCADE,
    comment_id  UUID        REFERENCES comments(id) ON DELETE CASCADE,
    entity_type VARCHAR(10) NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CHECK (
        (feed_id IS NOT NULL AND comment_id IS NULL AND entity_type = 'feed') OR
        (feed_id IS NULL AND comment_id IS NOT NULL AND entity_type = 'comment')
    )
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_feed_like
    ON likes (user_id, feed_id)
    WHERE feed_id IS NOT NULL;

CREATE UNIQUE INDEX IF NOT EXISTS unique_comment_like
    ON likes (user_id, comment_id)
    WHERE comment_id IS NOT NULL;

-- +goose Down
DROP INDEX IF EXISTS unique_feed_like;
DROP INDEX IF EXISTS unique_comment_like;
DROP TABLE IF EXISTS likes;
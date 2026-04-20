-- +goose Up
CREATE TABLE IF NOT EXISTS comments (
    id          UUID        PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    author_id   UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id     UUID        NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    parent_id   UUID        DEFAULT NULL REFERENCES comments(id) ON DELETE CASCADE,
    content     TEXT        NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_comments_feed_parent ON comments (feed_id, parent_id);
CREATE INDEX IF NOT EXISTS idx_comments_author_id   ON comments (author_id);

CREATE TRIGGER set_comment_updated_at
    BEFORE UPDATE ON comments
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();

-- +goose Down
DROP TRIGGER IF EXISTS set_comment_updated_at ON comments;
DROP TABLE IF EXISTS comments;
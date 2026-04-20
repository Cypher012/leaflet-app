-- +goose Up

CREATE TABLE IF NOT EXISTS feeds (
    id          UUID            PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id     UUID            NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title       VARCHAR(255)    NOT NULL,
    content     TEXT,
    feed_image  TEXT,
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_feeds_user_id        ON feeds (user_id);
CREATE INDEX IF NOT EXISTS idx_feeds_created_at_id  ON feeds (created_at DESC, id DESC);

CREATE TRIGGER set_feed_updated_at
    BEFORE UPDATE ON feeds
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();

-- +goose Down
DROP TRIGGER IF EXISTS set_feed_updated_at ON feeds;
DROP TABLE IF EXISTS feeds;
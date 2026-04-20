-- +goose Up
CREATE TABLE IF NOT EXISTS accounts (
    id                      UUID        PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id                 UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_id              VARCHAR(255) NOT NULL,
    provider_id             VARCHAR(50)  NOT NULL,
    access_token            TEXT,
    refresh_token           TEXT,
    access_token_expires_at TIMESTAMPTZ,
    password_hash           TEXT,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_user_provider UNIQUE (user_id, provider_id)
);

CREATE INDEX IF NOT EXISTS idx_account_provider ON accounts (provider_id, account_id);
CREATE INDEX IF NOT EXISTS idx_account_user_id  ON accounts (user_id);

CREATE TRIGGER set_account_updated_at
    BEFORE UPDATE ON accounts
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();

-- +goose Down
DROP TRIGGER IF EXISTS set_account_updated_at ON accounts;
DROP TABLE IF EXISTS accounts;
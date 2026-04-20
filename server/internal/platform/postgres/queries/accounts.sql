-- name: GetAccountByProvider :one
SELECT * FROM accounts
WHERE provider_id = $1
AND account_id = $2;

-- name: CreateAccount :one
INSERT INTO accounts (user_id, provider_id, account_id, access_token, refresh_token, password_hash, access_token_expires_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (provider_id, account_id) DO UPDATE SET
    access_token = EXCLUDED.access_token,
    refresh_token = EXCLUDED.refresh_token,
    access_token_expires_at = EXCLUDED.access_token_expires_at
RETURNING *;

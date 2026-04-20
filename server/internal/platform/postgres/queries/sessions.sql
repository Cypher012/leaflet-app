-- name: ListUserSessions :many
SELECT * FROM sessions
WHERE user_id = $1
AND expires_at > NOW();

-- name: CreateSession :one
INSERT INTO sessions (user_id, token, expires_at, ip_address, user_agent)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: RevokeSession :exec
UPDATE sessions
SET expires_at = NOW()
WHERE token = $1;

-- name: RevokeAllUserSessions :exec
UPDATE sessions
SET expires_at = NOW()
WHERE user_id = $1;

-- name: GetSessionByToken :one
SELECT * FROM sessions
WHERE token = $1
AND expires_at > NOW();


-- name: DeleteSessionByID :exec
DELETE FROM sessions
WHERE id = $1;

-- name: DeleteSessionByToken :exec
DELETE FROM sessions
WHERE token = $1;

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions
WHERE expires_at <= NOW();


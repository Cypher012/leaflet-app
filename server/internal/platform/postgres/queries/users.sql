-- name: CreateUser :one
INSERT INTO users (fullname, email, username, avatar_url)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = @email;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = @id;

-- name: GetUserIDByUsername :one
SELECT id FROM users
WHERE username = @username;


-- name: GetUserByToken :one
SELECT
    u.id,
    u.fullname,
    u.username,
    u.bio,
    u.avatar_url
FROM users u
JOIN sessions s ON s.user_id = u.id
WHERE s.token = $1
AND s.expires_at > NOW();


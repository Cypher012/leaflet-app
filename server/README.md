# Leaflet API (`server/`)

Go API for the Leaflet discussion platform.

## Tech Stack

- HTTP server: Echo v5
- DB: Postgres (`pgxpool`) + `sqlc`-generated query layer (`internal/platform/db`)
- Auth: OAuth (GitHub + Google via Goth) + cookie sessions stored in Postgres (`leaflet_sid`)
- Storage: Cloudflare R2 (S3-compatible) for direct uploads via presigned PUT URLs
- Docs: Swagger at `/docs/*`

## How The API Is Structured

- Entry: `cmd/server/main.go`
- Routing: `cmd/server/router.go` mounts modules under `/api`
- Modules live in `internal/module/*` and follow a consistent shape:
  - `route.go` defines endpoints + auth middleware
  - `handler.go` validates/parses HTTP and shapes responses
  - `service.go` holds business logic
  - `repository.go` calls the `sqlc` query layer
- Shared bits:
  - Auth middleware: `internal/shared/middleware/auth.go`
  - Response helpers: `internal/shared/response/*`
  - Cursor pagination helpers: `internal/shared/utils/*`

## Local Development

The default local setup expects the frontend to be served at `https://leaflet-dev.com` and the API at `https://api.leaflet-dev.com` (see repo root `Caddyfile`). The cookie is set with `Secure` + `SameSite=None`, so HTTPS matters.

### Run

```bash
make watch
```

If you don’t use hot reload:

```bash
make serve
```

## Configuration (Env Vars)

The API loads env vars in `cmd/server/main.go`:
- `APP_ENV=development` loads `.env` and `.env.development`
- Otherwise it loads the current environment

Required env vars (see `internal/platform/config/config.go`):

```txt
PORT
APP_ENV
DATABASE_URL
FRONTEND_URL
BACKEND_URL
COOKIE_DOMAIN
SESSION_SECRET

GITHUB_CLIENT_ID
GITHUB_CLIENT_SECRET
GITHUB_CALLBACK_URL

GOOGLE_CLIENT_ID
GOOGLE_CLIENT_SECRET
GOOGLE_CALLBACK_URL

R2_ACCOUNT_ID
R2_BUCKET
R2_ACCESS_KEY
R2_SECRET_KEY
R2_PUBLIC_URL
```

## Database

- Migrations: `internal/platform/postgres/migrations` (Goose)
- Queries: `internal/platform/postgres/queries` -> generated into `internal/platform/db` (sqlc)

Makefile targets assume you have `goose` installed and `DATABASE_URL` set:

```bash
make migrate-up
make migrate-status
```


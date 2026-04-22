# Leaflet

Leaflet is a full-stack discussion/forum-style app:
- **Frontend**: React + TanStack Start/Router/Query (`./web`)
- **Backend API**: Go (Echo) + Postgres + OAuth + R2 uploads (`./server`)

Local development is designed to run on HTTPS with two hostnames via Caddy:
- `https://leaflet-dev.com` (frontend)
- `https://api.leaflet-dev.com` (API)

This is important because the backend sets a **secure** session cookie (`leaflet_sid`, `SameSite=None; Secure`) for cross-subdomain auth.

## Architecture (High Level)

- `Caddyfile` terminates TLS (internal CA) and reverse-proxies:
  - `leaflet-dev.com -> localhost:3000` (Vite dev server)
  - `api.leaflet-dev.com -> localhost:8080` (Go API)
- The API serves routes under `/api/*` (see Swagger at `https://api.leaflet-dev.com/docs/` when running).
- The frontend talks to the API with credentials (cookies) enabled and uses React Query for caching/infinite scroll.
- Uploads are direct-to-R2 using a presigned URL returned by the API.

## Run Locally (Full Stack)

1. Add local DNS entries (usually `/etc/hosts`):

```txt
127.0.0.1 leaflet-dev.com
127.0.0.1 api.leaflet-dev.com
```

2. Start Caddy (from repo root):

```bash
caddy run
```

3. Start the backend API:

```bash
cd server
make watch
```

4. Start the frontend:

```bash
cd web
bun install
bun --bun run dev
```

Open `https://leaflet-dev.com`.

## Docs Per Package

- Backend: see `server/README.md`
- Frontend: see `web/README.md`


# Leaflet Web (`web/`)

Frontend for the Leaflet discussion platform.

## Tech Stack

- React + TanStack Start (SSR) + TanStack Router (file-based) + TanStack Query
- Tailwind CSS + shadcn/ui components
- Bundler/dev server: Vite
- Deploy target: Cloudflare Workers (Wrangler)

## Local Development

This frontend expects to be accessed via `https://leaflet-dev.com` with the API on `https://api.leaflet-dev.com` (repo root `Caddyfile`). This matters because auth uses a cross-subdomain secure cookie (`leaflet_sid`).

### Install & Run

```bash
bun install
bun --bun run dev
```

Vite runs on `http://localhost:3000`, but the intended dev URL is:
- `https://leaflet-dev.com` (Caddy reverse-proxies to `localhost:3000`)

### Environment Variables

See `.env.example`:

```txt
VITE_BASE_API_URL=
VITE_OAUTH_GITHUB_LOGIN=
VITE_OAUTH_GOOGLE_LOGIN=
```

Typical development values (see `.env.development`):

```txt
VITE_BASE_API_URL=https://api.leaflet-dev.com/api
VITE_OAUTH_GITHUB_LOGIN=https://api.leaflet-dev.com/api/auth/github
VITE_OAUTH_GOOGLE_LOGIN=https://api.leaflet-dev.com/api/auth/google
```

Notes:
- API calls are made with credentials enabled (`src/lib/axios.ts`), so the browser will send `leaflet_sid`.
- The dev server proxies `/api/*` requests to `https://api.leaflet-dev.com` (`vite.config.ts`).

## App Structure

- Routes: `src/routes/*` (file-based)
  - Home: `/_home/` shows the public feed list
  - Feed details: `/feeds/$feedId`
  - Profile: `/profile/$username/*`
  - Auth-gated pages: `src/routes/__protected/*` (redirects guests)
- Features: `src/features/*`
  - `feeds`: list, detail, comments, likes, infinite scroll
  - `create`: create feed (with optional image upload)
  - `profile`: overview/feeds/comments infinite lists
  - `auth`: OAuth entry points + login/signup modals
- API routes/constants: `src/lib/api-routes.ts`
- “Who am I” preload: `src/lib/auth.ts` (TanStack Start server function reads cookie and calls `/auth/me`)

## Upload Flow (R2)

1. Call `GET /api/upload/presign?type=feed|avatar|comment&content_type=...`
2. PUT the file directly to `upload_url`
3. Save `public_url` in the feed/comment payload

Implementation lives in `src/features/create/api/api.ts`.

## Scripts

```bash
bun --bun run dev
bun --bun run build
bun --bun run test
bun --bun run lint
bun --bun run format
bun --bun run check
```

### Deploy (Workers)

```bash
bun --bun run deploy
```

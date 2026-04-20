// src/lib/env.ts
interface ImportMetaEnv {
  readonly VITE_API_BASE_URL: string
  readonly VITE_OAUTH_GITHUB_LOGIN: string
  readonly VITE_OAUTH_GOOGLE_LOGIN: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}

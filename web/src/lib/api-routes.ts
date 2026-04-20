// src/lib/api-routes.ts
const BASE =
  import.meta.env.VITE_BASE_API_URL ?? process.env.VITE_BASE_API_URL ?? ''

export const API_ROUTES = {
  baseURL: BASE,
  auth: {
    me: `${BASE}/auth/me`,
    session: `${BASE}/auth/session`,
    github: import.meta.env.VITE_OAUTH_GITHUB_LOGIN,
    google: import.meta.env.VITE_OAUTH_GOOGLE_LOGIN,
    logout: `${BASE}/auth/logout`,
  },
  feeds: {
    list: `${BASE}/feeds`,
    create: `${BASE}/feeds`,
    detail: (id: string) => `${BASE}/feeds/${id}`,
  },
  comments: {
    list: (feedId: string) => `${BASE}/feeds/${feedId}/comments`,
    post_comment: (feedId: string) => `${BASE}/feeds/${feedId}/comments`,
    post_reply: (commentId: string) => `${BASE}/comments/${commentId}/replies`,
  },
  likes: {
    feed_like: (feed_id: string) => `${BASE}/feeds/${feed_id}/like`,
    comment_like: (comment_id: string) => `${BASE}/comments/${comment_id}/like`,
  },
  profile: {
    user: (username: string) => `${BASE}/profile/${username}`,
    overview: (username: string) => `${BASE}/profile/${username}/overview`,
    feeds: (username: string) => `${BASE}/profile/${username}/feeds`,
    comments: (username: string) => `${BASE}/profile/${username}/comments`,
  },
  upload: {
    presign: `${BASE}/upload/presign`,
  },
} as const

export const APP_ROUTES = {
  profile: '/profile',
  feeds: '/feeds',
}

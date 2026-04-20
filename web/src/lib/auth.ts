import { createServerFn } from '@tanstack/react-start'
import { getCookie } from '@tanstack/react-start/server'
import { API_ROUTES } from './api-routes'
import { queryOptions } from '@tanstack/react-query'
import type { UserResponse } from '#/types/user'

export const getMe = createServerFn({ method: 'GET' }).handler(
  async (): Promise<UserResponse | null> => {
    const token = getCookie('leaflet_sid')
    if (!token) return null

    try {
      const res = await fetch(API_ROUTES.auth.me, {
        headers: {
          Cookie: `leaflet_sid=${token}`,
        },
      })
      if (!res.ok) return null

      const json = await res.json()
      return json.data // unwrap here
    } catch (e) {
      console.log({ fetchError: String(e) })
      return null
    }
  },
)

export const getMeQueryOptions = queryOptions({
  queryKey: ['user'],
  queryFn: () => getMe(),
  staleTime: 5 * 60 * 1000,
})

import { apiClient } from '#/lib/axios'
// src/hooks/auth.useLogout.ts
import { useQueryClient } from '@tanstack/react-query'
import { useRouter } from '@tanstack/react-router'
import { API_ROUTES } from '#/lib/api-routes'

export function useLogout() {
  const queryClient = useQueryClient()
  const router = useRouter()

  async function logout() {
    try {
      await apiClient.post(API_ROUTES.auth.logout)
      await queryClient.invalidateQueries({ queryKey: ['user'] })
      queryClient.clear()
      await router.invalidate()
      await router.navigate({ to: '/' })
    } catch (error) {
      console.error('Logout failed:', error)
    }
  }

  return { logout }
}

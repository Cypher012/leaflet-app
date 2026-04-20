import { API_ROUTES } from '#/lib/api-routes'
import { apiClient } from '#/lib/axios'
import type { MessageResponse } from '#/types/global'
import { useMutation, useQueryClient } from '@tanstack/react-query'

const createFeed = async (payload: CreateCardReq): Promise<MessageResponse> => {
  const res = await apiClient.post<MessageResponse, CreateCardReq>(
    API_ROUTES.feeds.create,
    payload,
  )
  return res
}

export const useCreateFeed = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: createFeed,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['feeds'] })
    },
  })
}

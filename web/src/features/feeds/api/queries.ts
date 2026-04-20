import { infiniteQueryOptions, useQuery } from '@tanstack/react-query'
import { API_ROUTES } from '#/lib/api-routes'
import { apiClient } from '#/lib/axios'
import type { FeedComment, PublicFeed } from '../type'
import type { ApiResponse, PaginatedResponse } from '#/types/global'

const fetchFeeds = async (
  cursor: string,
): Promise<PaginatedResponse<PublicFeed>> => {
  const res = await apiClient.get<PaginatedResponse<PublicFeed>>(
    API_ROUTES.feeds.list,
    {
      params: cursor ? { cursor } : {},
    },
  )
  return res
}

const fetchFeed = async (feedId: string): Promise<PublicFeed> => {
  const res = await apiClient.get<ApiResponse<PublicFeed>>(
    API_ROUTES.feeds.detail(feedId),
  )
  return res.data
}

export const FeedsInfiniteQueryOptions = infiniteQueryOptions({
  queryKey: ['feeds'],
  queryFn: ({ pageParam }) => fetchFeeds(pageParam),
  getNextPageParam: (lastPage) =>
    lastPage.meta.has_next ? lastPage.meta.next_cursor : undefined,
  initialPageParam: '',
  staleTime: 1000 * 60 * 3, // feeds list goes stale faster than a single feed
  gcTime: 1000 * 60 * 10,
  refetchOnWindowFocus: false,
  refetchOnMount: false, // don't refetch when scrolling back to the list — preserves scroll position + optimistic updates
})

export const useFeedQuery = (feedId: string) =>
  useQuery({
    queryKey: ['feeds', feedId],
    queryFn: () => fetchFeed(feedId),
    refetchOnWindowFocus: false,
    staleTime: 1000 * 60 * 5, // don't refetch if data is less than 5 mins old
    gcTime: 1000 * 60 * 10, // keep in cache for 10 mins after unmount
    retry: 2, // retry failed requests twice before erroring
  })

const fetchComments = async (
  feedId: string,
  cursor: string,
): Promise<PaginatedResponse<FeedComment>> => {
  const res = await apiClient.get<PaginatedResponse<FeedComment>>(
    API_ROUTES.comments.list(feedId),
    {
      params: cursor ? { cursor } : {},
    },
  )
  return res
}

export const CommentsInfiniteQueryOptions = (feedId: string) =>
  infiniteQueryOptions({
    queryKey: ['comments', feedId],
    queryFn: ({ pageParam }) => fetchComments(feedId, pageParam),
    getNextPageParam: (lastPage) =>
      lastPage.meta.has_next ? lastPage.meta.next_cursor : undefined,
    initialPageParam: '',
    staleTime: 1000 * 60 * 3, // feeds list goes stale faster than a single feed
    gcTime: 1000 * 60 * 10,
    refetchOnWindowFocus: false,
    refetchOnMount: false,
  })

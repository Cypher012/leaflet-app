import { API_ROUTES } from "#/lib/api-routes";
import { apiClient } from "#/lib/axios";
import type { ApiResponse, PaginatedResponse } from "#/types/global";
import { infiniteQueryOptions, useQuery } from "@tanstack/react-query";
import type { ProfileCommentResponse, ProfileFeedsResponse, ProfileOverviewResponse, ProfileUserResponse} from '../types'


const fetchUserProfile = async(username: string): Promise<ProfileUserResponse> => {
    const res = await apiClient.get<ApiResponse<ProfileUserResponse>>(
        API_ROUTES.profile.user(username),
    )

    return res.data
}

export const useUserProfileQuery = (username: string) =>
  useQuery({
    queryKey: ['profile', username],
    queryFn: () => fetchUserProfile(username),
})

const fetchProfileOverview = async (
  username: string,
  cursor: string,
): Promise<PaginatedResponse<ProfileOverviewResponse>> => {
  const res = await apiClient.get<PaginatedResponse<ProfileOverviewResponse>>(
    API_ROUTES.profile.overview(username),
    {
      params: cursor ? { cursor } : {},
    },
  )
  return res
}

export const ProfileOverviewInfiniteQueryOptions = (username: string) => infiniteQueryOptions({
  queryKey: ['profile', username, "overview"],
  queryFn: ({ pageParam }) => fetchProfileOverview(username, pageParam),
  getNextPageParam: (lastPage) =>
    lastPage.meta.has_next ? lastPage.meta.next_cursor : undefined,
  initialPageParam: '',
})


const fetchProfileFeeds = async (
  username: string,
  cursor: string,
): Promise<PaginatedResponse<ProfileFeedsResponse>> => {
  const res = await apiClient.get<PaginatedResponse<ProfileFeedsResponse>>(
    API_ROUTES.profile.feeds(username),
    {
      params: cursor ? { cursor } : {},
    },
  )
  return res
}

export const ProfileFeedsInfiniteQueryOptions = (username: string) => infiniteQueryOptions({
  queryKey: ['profile', username, "feeds"],
  queryFn: ({ pageParam }) => fetchProfileFeeds(username,pageParam),
  getNextPageParam: (lastPage) =>
    lastPage.meta.has_next ? lastPage.meta.next_cursor : undefined,
  initialPageParam: '',
})

const fetchProfileComments = async (
  username: string,
  cursor: string,
): Promise<PaginatedResponse<ProfileCommentResponse>> => {
  const res = await apiClient.get<PaginatedResponse<ProfileCommentResponse>>(
    API_ROUTES.profile.comments(username),
    {
      params: cursor ? { cursor } : {},
    },
  )
  return res
}

export const ProfileCommentsInfiniteQueryOptions = (username: string) => infiniteQueryOptions({
  queryKey: ['profile', username, "comments"],
  queryFn: ({ pageParam }) => fetchProfileComments(username,pageParam),
  getNextPageParam: (lastPage) =>
    lastPage.meta.has_next ? lastPage.meta.next_cursor : undefined,
  initialPageParam: '',
})
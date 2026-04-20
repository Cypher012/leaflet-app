import { API_ROUTES } from '#/lib/api-routes'
import { apiClient } from '#/lib/axios'
import type { MessageResponse, PaginatedResponse } from '#/types/global'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import type { InfiniteData } from '@tanstack/react-query'
import type { CreateCommentReq, FeedComment, PublicFeed } from '../type'

const likeFeed = async (feedId: string): Promise<MessageResponse> => {
  const res = await apiClient.post<MessageResponse>(
    API_ROUTES.likes.feed_like(feedId),
  )
  return res
}

const likeComment = async (commentId: string): Promise<MessageResponse> => {
  const res = await apiClient.post<MessageResponse>(
    API_ROUTES.likes.comment_like(commentId),
  )
  return res
}

export const useLikeFeed = (feedId: string) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: () => likeFeed(feedId), // your API call

    onMutate: async () => {
      // Cancel any outgoing refetches so they don't overwrite our optimistic update
      await queryClient.cancelQueries({ queryKey: ['feeds'] })
      await queryClient.cancelQueries({ queryKey: ['feeds', feedId] })

      // Snapshot previous states for rollback
      const previousList = queryClient.getQueryData(['feeds'])
      const previousSingle = queryClient.getQueryData(['feeds', feedId])

      // Optimistic update for list (infinite query)
      queryClient.setQueryData(
        ['feeds'],
        (old: InfiniteData<PaginatedResponse<PublicFeed>> | undefined) => {
          if (!old) return old

          return {
            ...old,
            pages: old.pages.map((page: any) => ({
              ...page,
              data: page.data.map((feed: PublicFeed) =>
                feed.id === feedId
                  ? {
                      ...feed,
                      is_liked: !feed.is_liked,
                      stats: {
                        ...feed.stats,
                        likes: feed.is_liked
                          ? feed.stats.likes - 1
                          : feed.stats.likes + 1,
                      },
                    }
                  : feed,
              ),
            })),
          }
        },
      )

      // Optimistic update for single feed
      queryClient.setQueryData(
        ['feeds', feedId],
        (old: PublicFeed | undefined) => {
          if (!old) return old

          return {
            ...old,
            is_liked: !old.is_liked,
            stats: {
              ...old.stats,
              likes: old.is_liked ? old.stats.likes - 1 : old.stats.likes + 1,
            },
          }
        },
      )

      return { previousList, previousSingle }
    },

    onError: (err, _variables, context) => {
      // Rollback on error
      if (context?.previousList) {
        queryClient.setQueryData(['feeds'], context.previousList)
      }
      if (context?.previousSingle) {
        queryClient.setQueryData(['feeds', feedId], context.previousSingle)
      }

      // Optional: show toast error
      // toast.error("Failed to like feed. Please try again.");
    },

    onSettled: () => {
      // Invalidate to sync with server (safe even after optimistic update)
      queryClient.invalidateQueries({ queryKey: ['feeds'] })
      queryClient.invalidateQueries({ queryKey: ['feeds', feedId] })
    },
  })
}

export const useLikeComment = (feedId: string, commentId: string) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: () => likeComment(commentId),

    onMutate: async () => {
      // cancel any outgoing refetches so they don't overwrite out optimistic updates
      await queryClient.cancelQueries({ queryKey: ['comments', feedId] })

      // snapshot previous state for rollback
      const previousState = queryClient.getQueryData(['comments', feedId])

      // Optimistic update for list (infinite query)
      queryClient.setQueryData(
        ['comments', feedId],
        (old: InfiniteData<PaginatedResponse<FeedComment>> | undefined) => {
          if (!old) return old

          const localState: InfiniteData<PaginatedResponse<FeedComment>> = {
            ...old,
            pages: old.pages.map((page) => ({
              ...page,
              data: page.data.map((comment): FeedComment => {
                if (comment.id === commentId) {
                  return {
                    ...comment,
                    stats: {
                      ...comment.stats,
                      is_liked: !comment.stats.is_liked,
                      like_count: comment.stats.is_liked
                        ? comment.stats.like_count - 1
                        : comment.stats.like_count + 1,
                    },
                  }
                }

                const updatedReplies: FeedComment[] = comment.replies.map(
                  (reply): FeedComment =>
                    reply.id === commentId
                      ? {
                          ...reply,
                          stats: {
                            ...reply.stats,
                            is_liked: !reply.stats.is_liked,
                            like_count: reply.stats.is_liked
                              ? reply.stats.like_count - 1
                              : reply.stats.like_count + 1,
                          },
                        }
                      : reply,
                )

                return { ...comment, replies: updatedReplies }
              }),
            })),
          }

          return localState
        },
      )
      return { previousState }
    },

    onError: (err, _variables, context) => {
      // Rollback on error
      if (context?.previousState) {
        queryClient.setQueryData(['comment', feedId], context.previousState)
      }
    },

    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ['comments', feedId] })
    },
  })
}

const postComment = async (
  feedId: string,
  payload: CreateCommentReq,
): Promise<MessageResponse> => {
  const res = await apiClient.post<MessageResponse, CreateCommentReq>(
    API_ROUTES.comments.post_comment(feedId),
    payload,
  )
  return res
}

export const useCreateComment = (feedId: string, author: FeedComment["author"]) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (payload: CreateCommentReq) => postComment(feedId, payload),
    onMutate: async (payload) => {
      await queryClient.cancelQueries({ queryKey: ['comments', feedId] })

      const previousState = queryClient.getQueryData(['comments', feedId])

      const optimisticComment: FeedComment = {
        id: `temp-${Date.now()}`,
        parent_id: null,
        feed_id: feedId,
        content: payload.content,
        author,
        created_at: new Date().toISOString(),
        replies: [],
        stats: {
          is_liked: false,
          like_count: 0,
        },
      }

      queryClient.setQueryData(['comments', feedId], (old: InfiniteData<PaginatedResponse<FeedComment>> | undefined) => {
        if (!old) return old

        const localState: InfiniteData<PaginatedResponse<FeedComment>> = {
          ...old,
          pages: old.pages.map((page, index) =>
            index === 0
              ? {
                  ...page,
                  data: [optimisticComment, ...page.data],
                }
              : page
          ),
        }

        return localState
      })

      return {previousState}
    },

    onError: (_err, _variables, context) => {
      if (context?.previousState) {
        queryClient.setQueryData(['comments', feedId], context.previousState)
      }
    },

    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ['comments', feedId] })
      // just bump the comment count in the feed cache
      queryClient.setQueryData(['feeds', feedId], (old: PublicFeed | undefined) => {
        if (!old) return old
        return {
          ...old,
          stats: {
            ...old.stats,
            comments: old.stats.comments + 1,
          },
        }
      })
    },
  })
}

const postReply = async (
  commentId: string,
  payload: CreateCommentReq,
): Promise<MessageResponse> => {
  const res = await apiClient.post<MessageResponse, CreateCommentReq>(
    API_ROUTES.comments.post_reply(commentId),
    payload,
  )
  return res
}

export const useCreateReply = (feedId: string, commentId: string, author: FeedComment['author']) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (payload: CreateCommentReq) => postReply(commentId, payload),

    onMutate: async (payload) => {
      await queryClient.cancelQueries({ queryKey: ['comments', feedId] })

      const previousState = queryClient.getQueryData(['comments', feedId])

      const optimisticReply: FeedComment = {
        id: `temp-${Date.now()}`,
        parent_id: commentId,
        feed_id: feedId,
        content: payload.content,
        author,
        created_at: new Date().toISOString(),
        replies: [],
        stats: {
          is_liked: false,
          like_count: 0,
        },
      }
      queryClient.setQueryData(['comments', feedId], (old: InfiniteData<PaginatedResponse<FeedComment>> | undefined) => {
        if (!old) return old

        const localState: InfiniteData<PaginatedResponse<FeedComment>> = {
          ...old,
          pages: old.pages.map(page => ({
            ...page,
            data: page.data.map((comment): FeedComment =>
              comment.id === commentId
                ? { ...comment, replies: [...comment.replies, optimisticReply] }
                : comment
            ),
          })),
        }

        return localState
      })

      return { previousState }
    },

    onError: (_err, _variables, context) => {
      if (context?.previousState) {
        queryClient.setQueryData(['comments', feedId], context.previousState)
      }
    },

    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ['comments', feedId] })
      // just bump the comment count in the feed cache
      queryClient.setQueryData(['feeds', feedId], (old: PublicFeed | undefined) => {
        if (!old) return old
        return {
          ...old,
          stats: {
            ...old.stats,
            comments: old.stats.comments + 1,
          },
        }
      })
    },
  })
}

import { useLikeComment, useLikeFeed } from '#/features/feeds/api/mutation'

type LikePayloadProp =
  | { entity: 'feed'; feedId: string }
  | { entity: 'comment'; feedId: string; commentId: string }

export const useLike = (payload: LikePayloadProp) => {
  const { mutateAsync: likeFeed } = useLikeFeed(payload.feedId)
  const { mutateAsync: likeComment } = useLikeComment(
    payload.feedId,
    payload.entity === 'comment' ? payload.commentId : '',
  )

  const handleLike = async () => {
    if (payload.entity === 'feed') {
      await likeFeed()
    } else {
      await likeComment()
    }
  }

  return { handleLike }
}

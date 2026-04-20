import PulsarLoader from '#/components/web/pulsar-loader'
import { ProfileCommentsInfiniteQueryOptions } from '#/features/profile/api/queries'
import { useInfiniteScroll } from '#/hooks/useInfiiteScroll'
import { useInfiniteQuery } from '@tanstack/react-query'
import { CommentCard } from '#/features/profile/components/commentCard'
import OverviewPageSkeleton from '#/features/profile/components/skeleton/overview-page-skeleton'
import type { ProfileCommentResponse } from '#/features/profile/types'

const CommentPage = ({ username }: { username: string }) => {
  const query = useInfiniteQuery(ProfileCommentsInfiniteQueryOptions(username))

  const {
    items: comments,
    bottomRef,
    isLoading,
    isFetchingNextPage,
  } = useInfiniteScroll<ProfileCommentResponse>(query)

  if (isLoading) {
      return <OverviewPageSkeleton />
   }

  return (
    <div className="mt-10 space-y-5">
      <div className="space-y-6">
        {
            comments.map((comment) => (
                <CommentCard
                    key={comment.id}
                    id={comment.id}
                    feedId={comment.feed_id}
                    postTitle={comment.title}
                    commentBody={comment.content}
                    timestamp={comment.created_at}
                    is_liked={comment.stats.is_liked}
                    like_count={comment.stats.like_count}
                />
            ))
        }
        <div ref={bottomRef} />
        {isFetchingNextPage && <PulsarLoader />}
      </div>
    </div>
  )
}

export default CommentPage
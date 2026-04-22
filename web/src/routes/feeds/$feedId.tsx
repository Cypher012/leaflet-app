import { createFileRoute, useRouteContext, useRouter } from '@tanstack/react-router'
import { useFeedQuery } from '#/features/feeds/api/queries'
import { useLike } from '#/hooks/useLike'
import { FeedDetailSkeleton } from '#/features/feeds/components/feed-details-skeleton'
import { FeedDetailPage } from '#/features/feeds/components/page/feedDetailPage'
import z from 'zod'

export const Route = createFileRoute('/feeds/$feedId')({
  validateSearch: z.object({
    commentId: z.string().optional(),
  }),
  component: RouteComponent,
})

function RouteComponent() {
  const { user } = useRouteContext({ from: '__root__' })
  const { feedId } = Route.useParams()
  const { commentId } = Route.useSearch()
  const { data: feed, isLoading } = useFeedQuery(feedId)
  const { handleLike } = useLike({ entity: 'feed', feedId })
  const router = useRouter()

  if (isLoading) {
    return <FeedDetailSkeleton />
  }

  if (!feed) return null

  return (
    <FeedDetailPage
    feed={feed}
    user={user}
    feedId={feedId}
    commentId={commentId}
    onBack={() => router.history.back()}
    onLike={handleLike}
    />
  )
}
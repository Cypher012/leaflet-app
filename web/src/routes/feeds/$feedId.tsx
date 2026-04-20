import { Avatar, AvatarFallback, AvatarImage } from '#/components/ui/avatar'
import LikeButton from '#/components/ui/likeButton'
import ArchiveDiscussion from '#/features/feeds/components/archiveDiscussion'
import { useFeedQuery } from '#/features/feeds/api/queries'
import { useLike } from '#/hooks/useLike'
import { formatRelativeTime } from '#/lib/time'
import {
  createFileRoute,
  useRouteContext,
  useRouter,
} from '@tanstack/react-router'
import { ArrowLeft, Dot, MessageSquare } from 'lucide-react'
import { Button } from '#/components/ui/button'
import { FeedDetailSkeleton } from '#/features/feeds/components/feed-details-skeleton'

export const Route = createFileRoute('/feeds/$feedId')({
  component: RouteComponent,
})

function RouteComponent() {
  const { user } = useRouteContext({ from: '__root__' })
  const { feedId } = Route.useParams()
  const { data: feed, isLoading, isFetching } = useFeedQuery(feedId)
  const { handleLike } = useLike({ entity: 'feed', feedId })

  const router = useRouter()

  if (isLoading || isFetching) {
    return <FeedDetailSkeleton />
  }

  if (!feed) {
    return null
  }

  const handleScrollToComments = () => {
    document
      .getElementById('comment-box')
      ?.scrollIntoView({ behavior: 'smooth' })
  }

  return (
    <div className="relative section-container max-w-3xl space-y-10">
      <div className="absolute top-14 left-0">
        <Button
          onClick={() => router.history.back()}
          variant={'secondary'}
          className="size-10"
        >
          <ArrowLeft className="size-5" />
        </Button>
      </div>
      <div className="p-5 space-y-7">
        <span className="flex  text-sm items-center text-muted-foreground">
          <Dot /> Published {formatRelativeTime(feed.created_at)}
        </span>
        <div className="space-y-4">
          <h1 className="font-bold text-3xl text-foreground ">{feed.title}</h1>
          <div className="flex gap-x-5 items-center">
            <Avatar className="size-10">
              <AvatarImage src={feed.author.avatar_url} alt="avatar image" />
              <AvatarFallback>{feed.author.fullname[0]}</AvatarFallback>
            </Avatar>
            <p className="font-semibold text-foreground">
              {feed.author.fullname}
            </p>
          </div>
        </div>
        {feed.feed_image && (
          <div className="relative aspect-video max-w-full  mx-auto rounded-2xl overflow-hidden">
            <img
              src={feed.feed_image}
              alt={feed.title}
              className="absolute inset-0 h-full w-full object-cover"
            />
          </div>
        )}
        <div className="">
          <p className="text-muted-foreground leading-relaxed text-sm">
            {feed.content}
          </p>
        </div>

        <div className="flex items-center gap-6 text-sm text-muted-foreground">
          <LikeButton
            handleLike={handleLike}
            isLiked={feed.is_liked}
            likes={feed.stats.likes}
          />

          <button
            onClick={handleScrollToComments}
            className="flex items-center gap-2"
          >
            <MessageSquare className="w-4 h-4" />
            <span>{feed.stats.comments}</span>
          </button>
        </div>
      </div>
      <div id="comment-box" className="">
        <ArchiveDiscussion user={user} feedId={feedId} />
      </div>
    </div>
  )
}

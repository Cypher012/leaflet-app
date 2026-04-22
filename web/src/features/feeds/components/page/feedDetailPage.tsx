import { Avatar, AvatarFallback, AvatarImage } from '#/components/ui/avatar'
import LikeButton from '#/components/ui/likeButton'
import ArchiveDiscussion from '#/features/feeds/components/archiveDiscussion'
import { formatRelativeTime } from '#/lib/time'
import { ArrowLeft, Dot, MessageSquare } from 'lucide-react'
import { Button } from '#/components/ui/button'

type Props = {
  feed: any
  user: any
  feedId: string
  commentId: string | undefined
  onBack: () => void
  onLike: () => Promise<void>
}

export function FeedDetailPage({
  feed,
  user,
  feedId,
  commentId,
  onBack,
  onLike,
}: Props) {
  const handleScrollToComments = () => {
    document
      .getElementById('comment-box')
      ?.scrollIntoView({ behavior: 'smooth' })
  }

  return (
    <div className="relative section-container max-w-3xl">
      <div className="absolute top-14 -left-5 -lg:left-5 lg:flex hidden">
        <Button onClick={onBack} variant="secondary" className="size-10">
          <ArrowLeft className="size-5" />
        </Button>
      </div>

      {/* Mobile back button */}
      <div className="lg:hidden flex mt-10">
        <Button onClick={onBack} variant="secondary" className="size-10">
          <ArrowLeft className="size-5" />
        </Button>
      </div>

      <div className="space-y-10 mt-10">
        <div className="p-5 space-y-7 relative">
          <span className="text-muted-foreground md:text-sm text-xs flex items-center absolute top-0 left-3.5">
            <Dot /> Published {formatRelativeTime(feed.created_at)}
          </span>

          <div className="space-y-4 mt-5">
            <h1 className="font-bold md:text-3xl text-2xl text-foreground">
              {feed.title}
            </h1>

            <div className="flex md:gap-x-5 gap-x-3 items-center">
              <Avatar className="md:size-10 size-8">
                <AvatarImage src={feed.author.avatar_url} />
                <AvatarFallback>
                  {feed.author.fullname[0]}
                </AvatarFallback>
              </Avatar>
              <p className="font-semibold md:text-base text-sm text-foreground">
                {feed.author.fullname}
              </p>
            </div>
          </div>

          {feed.feed_image && (
            <div className="relative aspect-video max-w-full mx-auto rounded-2xl overflow-hidden">
              <img
                src={feed.feed_image}
                alt={feed.title}
                className="absolute inset-0 h-full w-full object-cover"
              />
            </div>
          )}

          <p className="text-muted-foreground leading-relaxed md:text-sm text-xs">
            {feed.content}
          </p>

          <div className="flex items-center gap-6 text-sm text-muted-foreground">
            <LikeButton
              handleLike={onLike}
              isLiked={feed.is_liked}
              likes={feed.stats.likes}
            />

            <button
              onClick={handleScrollToComments}
              className="flex items-center gap-2"
            >
              <MessageSquare className="md:size-4 size-3.5" />
              <span>{feed.stats.comments}</span>
            </button>
          </div>
        </div>

        <div id="comment-box">
          <ArchiveDiscussion user={user} commentId={commentId} feedId={feedId} />
        </div>
      </div>

    </div>
  )
}
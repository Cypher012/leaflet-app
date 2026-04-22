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
    <div className="section-container max-w-3xl space-y-10">
      {/* Back button — inline on mobile, absolute on md+ */}
      <div className="md:absolute md:top-14 md:left-0 pt-4 md:pt-0">
        <Button onClick={onBack} variant="secondary" className="size-10">
          <ArrowLeft className="size-5" />
        </Button>
      </div>

      <div className="p-4 md:p-5 space-y-7">
        <span className="flex text-sm items-center text-muted-foreground">
          <Dot /> Published {formatRelativeTime(feed.created_at)}
        </span>

        <div className="space-y-4">
          <h1 className="font-bold text-2xl md:text-3xl text-foreground">
            {feed.title}
          </h1>

          <div className="flex gap-x-3 md:gap-x-5 items-center">
            <Avatar className="size-8 md:size-10 shrink-0">
              <AvatarImage src={feed.author.avatar_url} />
              <AvatarFallback>
                {feed.author.fullname[0]}
              </AvatarFallback>
            </Avatar>
            <p className="font-semibold text-foreground truncate">
              {feed.author.fullname}
            </p>
          </div>
        </div>

        {feed.feed_image && (
          <div className="relative aspect-video w-full mx-auto rounded-2xl overflow-hidden">
            <img
              src={feed.feed_image}
              alt={feed.title}
              className="absolute inset-0 h-full w-full object-cover"
            />
          </div>
        )}

        <p className="text-muted-foreground leading-relaxed text-sm">
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
            <MessageSquare className="w-4 h-4" />
            <span>{feed.stats.comments}</span>
          </button>
        </div>
      </div>

      <div id="comment-box">
        <ArchiveDiscussion user={user} commentId={commentId} feedId={feedId} />
      </div>
    </div>
  )
}
import { MessageSquare } from 'lucide-react'
import { Link } from '@tanstack/react-router'
import { useLike } from '#/hooks/useLike'
import type { PublicFeed } from '../type'
import { formatRelativeTime } from '#/lib/time'
import { Avatar, AvatarFallback, AvatarImage } from '#/components/ui/avatar'
import LikeButton from '#/components/ui/likeButton'

export default function FeedCard({
  id: feedId,
  title,
  content,
  author,
  is_liked,
  created_at,
  feed_image,
  stats,
}: PublicFeed) {
  const { handleLike } = useLike({ entity: 'feed', feedId })

  return (
    <Link
      to="/feeds/$feedId"
      params={{ feedId: `${feedId}` }}
      className="rounded-4xl bg-muted overflow-hidden block"
    >
      <div className="space-y-4 bg-card p-4 md:p-6">
        {/* Title */}
        <h2 className="text-base md:text-lg font-semibold leading-snug">
          {title}
        </h2>

        {/* Image (optional) */}
        {feed_image && (
          <div className="relative aspect-video w-full md:max-w-xl md:mx-auto rounded-2xl overflow-hidden">
            <img
              src={feed_image}
              alt={title}
              loading="lazy"
              className="absolute inset-0 h-full w-full object-cover"
            />
          </div>
        )}

        {/* Content */}
        <p className="md:text-sm text-xs text-muted-foreground leading-relaxed line-clamp-3">
          {content}
        </p>

        {/* Stats */}
        <div className="flex items-center gap-4 md:text-sm text-xs text-muted-foreground">
          <LikeButton
            handleLike={handleLike}
            isLiked={is_liked}
            likes={stats.likes}
          />
          <div className="flex items-center gap-2">
            <MessageSquare className="md:size-4 size-3.5" />
            <span>{stats.comments}</span>
          </div>
        </div>
      </div>

      {/* Footer */}
      <div className="flex items-center justify-between gap-2 px-4 py-3 md:px-6 md:py-4">
        <div className="flex items-center gap-2 min-w-0">
          <Avatar className="w-7 h-7 md:w-8 md:h-8 shrink-0">
            <AvatarImage src={author.avatar_url} alt={author.fullname} />
            <AvatarFallback>{author.fullname[0]}</AvatarFallback>
          </Avatar>
          <span className="md:text-sm text-xs font-medium truncate">{author.fullname}</span>
        </div>

        <span className="md:text-xs text-[10px] text-muted-foreground shrink-0 whitespace-nowrap">
          {formatRelativeTime(created_at)}
        </span>
      </div>
    </Link>
  )
}
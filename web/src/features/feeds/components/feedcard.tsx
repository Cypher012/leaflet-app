import { Leaf, MessageSquare } from 'lucide-react'
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
      className={`rounded-4xl bg-muted overflow-hidden`}
    >
      <div className="space-y-6 bg-card p-6 ">
        {/* Title */}
        <h2 className="text-lg font-semibold">{title}</h2>

        {/* Image (optional) */}
        {feed_image && (
          <div className="relative aspect-video max-w-xl mx-auto rounded-2xl overflow-hidden">
            <img
              src={feed_image}
              alt={title}
              loading="lazy"
              className="absolute inset-0 h-full w-full object-cover"
            />
          </div>
        )}

        {/* Content */}
        <p className="text-sm text-muted-foreground leading-relaxed">
          {content}
        </p>

        {/* Stats */}
        <div className="flex items-center gap-6 text-sm text-muted-foreground">
          {/* <LikeButton /> */}
          <LikeButton
            handleLike={handleLike}
            isLiked={is_liked}
            likes={stats.likes}
          />
          <div className="flex items-center gap-2">
            <MessageSquare className="w-4 h-4" />
            <span>{stats.comments}</span>
          </div>
        </div>
      </div>

      {/* Footer */}
      <div className="flex items-center justify-between p-6">
        <div className="flex items-center gap-3">
          <Avatar className="w-8 h-8">
            <AvatarImage src={author.avatar_url} alt={author.fullname} />
            <AvatarFallback>{author.fullname[0]}</AvatarFallback>
          </Avatar>
          <span className="text-sm font-medium">{author.fullname}</span>
        </div>

        <span className="text-xs text-muted-foreground">
          {formatRelativeTime(created_at)}
        </span>
      </div>
    </Link>
  )
}

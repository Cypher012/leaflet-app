import { Dot, Heart, MessageSquare } from 'lucide-react'
import type { FeedCardProps } from '../types'
import { formatRelativeTime } from '#/lib/time'
import { Link } from '@tanstack/react-router'

const FeedCardProfile = ({
  id,
  title,
  content,
  image,
  createdAt,
  stats,
}: FeedCardProps) => {
  return (
    <Link to='/feeds/$feedId' params={{feedId: id}} className="bg-card text-card-foreground rounded-2xl p-5 flex gap-4 w-full">
      <div className="flex flex-col flex-1 gap-3 relative">
        <span className="text-muted-foreground text-xs flex items-center absolute top-0 -left-1.75">
          <Dot />
          {formatRelativeTime(createdAt)}
        </span>

        <h2 className="font-bold text-foreground text-xl leading-snug mt-10">
          {title}
        </h2>

        <p className="text-sm text-muted-foreground leading-relaxed flex-1 line-clamp-3">
          {content}
        </p>

        <div className="flex items-center gap-4 text-muted-foreground">
          <button className="flex items-center gap-1.5 text-sm hover:text-foreground transition-colors">
            <Heart className="size-4" />
            <span>{stats.like_count}</span>
          </button>
          <button className="flex items-center gap-1.5 text-sm hover:text-foreground transition-colors">
            <MessageSquare className="size-4" />
            <span>{stats.comment_count}</span>
          </button>
        </div>
      </div>

      {image && (
        <div className="shrink-0 w-36 h-36  rounded-md overflow-hidden">
          <img src={image} alt={title} className="w-full h-full object-cover" />
        </div>
      )}
    </Link>
  )
}

export default  FeedCardProfile

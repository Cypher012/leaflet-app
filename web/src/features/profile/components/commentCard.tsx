import { MessageSquare, ThumbsUp } from 'lucide-react'
import type { CommentCardProps } from '../types'
import { formatRelativeTime } from '#/lib/time'

export const CommentCard = ({
  postTitle,
  commentBody,
  timestamp,
  likes,
}: CommentCardProps) => {
  return (
    <div className="max-w-full p-5 bg-secondary/20 rounded-2xl">
      <div className="space-y-6">
        {/* Header Section */}
        <div className="flex items-center gap-3">
          <div className="bg-primary p-1.5 rounded-sm">
            <MessageSquare className="w-4 h-4 text-white fill-white" />
          </div>
          <p className="text-[10px] font-bold tracking-widest text-accent uppercase">
            Commented on{' '}
            <span className="text-secondary-foreground ml-1">{postTitle}</span>
          </p>
        </div>

        {/* Comment Body with Accent Border */}
        <div className="relative pl-6">
          <div className="absolute left-0 top-0 bottom-0 w-[3px] bg-[#A7F3D0] rounded-full" />
          <p className="text-lg italic text-muted-foreground leading-relaxed font-medium">
            &quot;{commentBody}&quot;
          </p>
        </div>

        {/* Footer Section */}
        <div className="flex justify-between items-center pt-2">
          <span className="text-xs font-medium text-[#7A8C83]">
            {formatRelativeTime(timestamp)}
          </span>
          <div className="flex items-center gap-2 text-[#7A8C83]">
            <ThumbsUp className="w-4 h-4 fill-current opacity-70" />
            <span className="text-sm font-semibold">{likes}</span>
          </div>
        </div>
      </div>
    </div>
  )
}

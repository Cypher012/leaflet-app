import {  Leaf, MessageSquare, ThumbsUp } from 'lucide-react'
import type { CommentCardProps } from '../types'
import { formatRelativeTime } from '#/lib/time'
import { Link } from '@tanstack/react-router'
import { useLike } from '#/hooks/useLike'
import { cn } from '#/lib/utils'

export const CommentCard = ({
  id,
  feedId,
  postTitle,
  commentBody,
  timestamp,
  like_count,
  is_liked
}: CommentCardProps) => {

  const {handleLike} = useLike({entity:"comment", feedId, commentId: id})

  return (
    <Link to="/feeds/$feedId" params={{feedId: feedId}} search={{commentId: id}} className="block max-w-full p-5 bg-secondary/20 rounded-2xl">
      <div className="space-y-6">
        {/* Header Section */}
        <div className="flex items-center gap-3">
          <div className="bg-primary p-1.5 rounded-sm">
            <MessageSquare className="w-4 h-4 text-white fill-white" />
          </div>
          <p className="md:text-[10px] text-[9px] font-bold tracking-widest text-accent uppercase">
            Commented on{' '}
            <span className="text-secondary-foreground ml-1">&quot;{postTitle}&quot;</span>
          </p>
        </div>

        {/* Comment Body with Accent Border */}
        <div className="relative pl-4">
          <div className="absolute left-0 top-0 bottom-0 w-[3px] bg-[#A7F3D0] rounded-full" />
          <p className="md:text-lg text-sm italic text-muted-foreground leading-relaxed font-medium">
            {commentBody}
          </p>
        </div>

        {/* Footer Section */}
        <div className="flex justify-between items-center pt-2">
          <span className="md:text-xs text-[10px] font-medium text-[#7A8C83]">
            {formatRelativeTime(timestamp)}
          </span>
          <div className="flex items-center text-xs gap-2 text-[#7A8C83]">
           <Leaf className={cn('size-4 md:size-3.5', is_liked ? 'text-primary fill-primary' : '')} />
            <span>{like_count}</span>
          </div>
        </div>
      </div>
    </Link>
  )
}

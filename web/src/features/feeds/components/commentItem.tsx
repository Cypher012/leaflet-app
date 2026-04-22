import { useEffect, useState } from 'react'
import { Avatar, AvatarFallback, AvatarImage } from '#/components/ui/avatar'
import { Button } from '#/components/ui/button'
import { Textarea } from '#/components/ui/textarea'
import { useLike } from '#/hooks/useLike'
import { formatRelativeTime } from '#/lib/time'
import type { FeedComment } from '../type'
import { useCreateReply } from '../api/mutation'
import LikeButton from '#/components/ui/likeButton'
import type { UserResponse } from '#/types/user'
import { cn } from '#/lib/utils'

export const CommentItem = ({
  user,
  comment,
  feedId,
  isReply = false,
}: {
  user: UserResponse | null
  comment: FeedComment
  feedId: string
  isReply?: boolean
}) => {
  const [collapsed, setCollapsed] = useState(false)
  const [replyOpen, setReplyOpen] = useState(false)
  const [replyValue, setReplyValue] = useState('')
  const { handleLike } = useLike({
    entity: 'comment',
    feedId,
    commentId: comment.id,
  })

  const { mutate: postReply, isPending } = useCreateReply(
    feedId,
    comment.id,
    user as FeedComment['author'],
  )

  const handleReply = () => {
    if (!replyValue.trim()) return
    setReplyValue('')
    setReplyOpen(false)
    postReply({ content: replyValue })
  }

  const isOptimistic = comment.id.startsWith('temp-')

  return (
    <div
      id={`comment-${comment.id}`}
      className={cn('flex gap-2 md:gap-3', isOptimistic && 'opacity-50 pointer-events-none')}
    >
      <div className="flex flex-col items-center">
        <Avatar className={isReply ? 'size-7 md:size-8' : 'size-8 md:size-10'}>
          <AvatarImage
            src={comment.author.avatar_url}
            alt={comment.author.fullname}
          />
          <AvatarFallback>{comment.author.fullname[0]}</AvatarFallback>
        </Avatar>
        {!isReply && comment.replies.length > 0 && !collapsed && (
          <div className="w-px flex-1 bg-border mt-2 mb-4" />
        )}
      </div>

      <div className="flex-1 min-w-0 pb-4 space-y-3">
        <div className="flex items-center gap-2 flex-wrap">
          <span className="font-semibold text-sm text-foreground truncate">
            {comment.author.fullname}
          </span>
          <span className="text-xs text-muted-foreground shrink-0">
            • {formatRelativeTime(comment.created_at)}
          </span>
        </div>

        <p className="text-sm text-foreground leading-relaxed wrap-break-word">
          {comment.content}
        </p>

        <div className="flex items-center gap-3 md:gap-4 flex-wrap">
          {comment.replies.length > 0 && (
            <button
              onClick={() => setCollapsed(!collapsed)}
              className="text-xs font-semibold tracking-widest text-muted-foreground hover:text-foreground transition-colors uppercase"
            >
              {collapsed ? 'Expand' : 'Collapse'}
            </button>
          )}
          <button
            onClick={() => setReplyOpen(!replyOpen)}
            className="text-xs font-semibold tracking-widest text-muted-foreground hover:text-foreground transition-colors uppercase"
          >
            {replyOpen ? 'Cancel' : 'Reply'}
          </button>
          <LikeButton
            handleLike={handleLike}
            isLiked={comment.stats.is_liked}
            likes={comment.stats.like_count}
          />
        </div>

        {replyOpen && (
          <div className="flex flex-col gap-2 mt-3">
            <Textarea
              value={replyValue}
              onChange={(e) => setReplyValue(e.target.value)}
              placeholder={`Reply to ${comment.author.fullname}...`}
              className="min-h-20 bg-muted border-0 resize-none text-sm p-3"
              autoFocus
            />
            <div className="flex justify-end gap-2">
              <Button
                variant="ghost"
                size="sm"
                onClick={() => {
                  setReplyOpen(false)
                  setReplyValue('')
                }}
              >
                Cancel
              </Button>
              <Button
                onClick={handleReply}
                disabled={isPending || !replyValue.trim()}
                size="sm"
                className="rounded-full px-4 bg-[#1a3a2a] hover:bg-[#1a3a2a]/90 text-white"
              >
                {isPending ? 'Posting...' : 'Post Reply'}
              </Button>
            </div>
          </div>
        )}

        {!collapsed && comment.replies.length > 0 && (
          <div className="mt-4 flex flex-col gap-4">
            {comment.replies.map((reply) => (
              <CommentItem
                key={reply.id}
                comment={reply}
                feedId={feedId}
                user={user}
                isReply
              />
            ))}
          </div>
        )}
      </div>
    </div>
  )
}

export default CommentItem
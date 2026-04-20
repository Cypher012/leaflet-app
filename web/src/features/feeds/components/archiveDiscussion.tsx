import { useEffect, useRef, useState } from 'react'
import { Button } from '#/components/ui/button'
import { Textarea } from '#/components/ui/textarea'
import type { UserResponse } from '#/types/user'
import CommentItem from './commentItem'
import { useInfiniteScroll } from '#/hooks/useInfiiteScroll'
import { useInfiniteQuery } from '@tanstack/react-query'
import type { FeedComment } from '../type'
import { CommentsInfiniteQueryOptions } from '../api/queries'
import { useElementVisibility } from '@reactuses/core'
import { useCreateComment } from '../api/mutation'
import PulsarLoader from '#/components/web/pulsar-loader'

const ArchiveDiscussion = ({
  user,
  feedId,
  commentId,
}: {
  user: UserResponse | null
  feedId: string
  commentId?:string
}) => {
  const [value, setValue] = useState('')
  const [shouldFetch, setShouldFetch] = useState(false)
  const [done, setDone] = useState(false)
  const { mutateAsync: postComment, isPending } = useCreateComment(feedId, user as FeedComment["author"])
  const ref = useRef<HTMLDivElement | null>(null)
  
  
  const query = useInfiniteQuery({
    ...CommentsInfiniteQueryOptions(feedId),
    enabled: shouldFetch,
  })

  const {
    items: comments,
    bottomRef,
    isFetchingNextPage,
    isLoading,
  } = useInfiniteScroll<FeedComment>(query)

  const [isVisible] = useElementVisibility(ref)

  useEffect(() => {
    if (commentId) {
      setShouldFetch(true)
      return
    }

    if (isVisible && !shouldFetch) {
      setShouldFetch(true)
    }
  }, [isVisible, commentId, shouldFetch])
  
  useEffect(() => {
    if (!commentId || done) return

    const el = document.getElementById(`comment-${commentId}`)

    if (el) {
      el.scrollIntoView({ block: 'center' })
      setDone(true)
      return
    }

    if (query.hasNextPage && !query.isFetchingNextPage) {
      query.fetchNextPage()
    }
  }, [commentId, comments, query.hasNextPage, query.isFetchingNextPage, done])
  
  const handlePostComment = async () => {
      if (!value.trim()) return
      setValue('')
      postComment({ content: value })
  }


  return (
    <div className="p-5">
      <h2 className="text-2xl font-bold text-foreground mb-6">
        Archive Discussion
      </h2>

      <div ref={ref} className="flex flex-col gap-3">
        <Textarea
          value={value}
          onChange={(e) => setValue(e.target.value)}
          placeholder="Add your observation..."
          className="min-h-28 bg-muted border-0 resize-none text-sm p-5"
        />
        <div className="flex justify-end">
          <Button
            onClick={handlePostComment}
            disabled={isPending || !value.trim()}
            className="rounded-full px-6 bg-[#1a3a2a] hover:bg-[#1a3a2a]/90 text-white"
          >
            Post comment
          </Button>
        </div>
      </div>

      <div className="mt-8 flex flex-col gap-2">
        {isLoading ? (
          <div className="flex justify-center">
            <PulsarLoader />
          </div>
        ) : comments.length === 0 ? (
          <p className="text-muted-foreground text-sm">
            No comments yet. Be the first!
          </p>
        ) : (
          comments.map((comment) => (
            <CommentItem key={comment.id} comment={comment} feedId={feedId} user={user} />
          ))
        )}
        <div ref={bottomRef} />
        {isFetchingNextPage && (
          <p className="text-muted-foreground text-sm text-center">
            Loading more...
          </p>
        )}
      </div>
    </div>
  )
}

export default ArchiveDiscussion

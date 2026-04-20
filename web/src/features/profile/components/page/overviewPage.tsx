import PulsarLoader from '#/components/web/pulsar-loader'
import { ProfileOverviewInfiniteQueryOptions } from '#/features/profile/api/queries'
import { useInfiniteScroll } from '#/hooks/useInfiiteScroll'
import { useInfiniteQuery } from '@tanstack/react-query'
import { CommentCard } from '#/features/profile/components/commentCard'
import FeedCardProfile from '#/features/profile/components/feedcardProfile'
import OverviewPageSkeleton from '#/features/profile/components/skeleton/overview-page-skeleton'
import type { ProfileOverviewResponse } from '#/features/profile/types'
import { useEffect } from 'react'

const OverviewPage = ({ username }: { username: string }) => {
  const query = useInfiniteQuery(ProfileOverviewInfiniteQueryOptions(username))

  const {
    items: overview,
    bottomRef,
    isLoading,
    isFetchingNextPage,
  } = useInfiniteScroll<ProfileOverviewResponse>(query)

  useEffect(() => {
    console.log(bottomRef.current)
    if(bottomRef.current) {
        console.log("visible")
    }else{
        console.log("not visible")
    }

  }, [bottomRef.current])

  
  if (isLoading) {
    return <OverviewPageSkeleton />
  }





  return (
    <div className="mt-10 space-y-5">
      <div className="space-y-6">
        {overview.map((item) => {
          if (item.type === 'feed') {
            return (
              <FeedCardProfile
                key={item.id}
                id={item.id}
                title={item.title}
                content={item.content}
                createdAt={item.created_at}
                image={item.feed_image}
                stats={item.stats}
              />
            )
          }

          // eslint-disable-next-line @typescript-eslint/no-unnecessary-condition
          if (item.type === 'comment') {
            return (
              <CommentCard
                key={item.id}
                id={item.id}
                feedId={item.comment_feed_id}
                postTitle={item.parent_feed_title}
                commentBody={item.comment_body}
                timestamp={item.created_at}
                is_liked={item.stats.is_liked}
                like_count={item.stats.like_count}
              />
            )
          }

          return null
        })}

        <div ref={bottomRef} />

        <div className="flex justify-center items-center">
            {isFetchingNextPage && <PulsarLoader />}
        </div>
      </div>
    </div>
  )
}

export default OverviewPage
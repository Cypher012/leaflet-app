import { ProfileFeedsInfiniteQueryOptions } from '#/features/profile/api/queries'
import { useInfiniteScroll } from '#/hooks/useInfiiteScroll'
import { useInfiniteQuery } from '@tanstack/react-query'
import type { ProfileFeedsResponse } from '#/features/profile/types'
import OverviewPageSkeleton from '../skeleton/overview-page-skeleton'
import PulsarLoader from '#/components/web/pulsar-loader'
import FeedCardProfile from '../feedcardProfile'

const FeedPage = ({ username }: { username: string }) => {
  const query = useInfiniteQuery(ProfileFeedsInfiniteQueryOptions(username))

  const {
    items: feeds,
    bottomRef,
    isLoading,
    isFetchingNextPage,
  } = useInfiniteScroll<ProfileFeedsResponse>(query)

  if (isLoading) {
      return <OverviewPageSkeleton />
   }

  return (
    <div className="mt-10 space-y-5">
      <div className="space-y-6">
        {
            feeds.map((feed) => (
                <FeedCardProfile
                key={feed.id}
                id={feed.id}
                title={feed.title}
                content={feed.content}
                createdAt={feed.created_at}
                image={feed.feed_image}
                stats={feed.stats}
                />
            ))
        }
        <div ref={bottomRef} />
        {isFetchingNextPage && <PulsarLoader />}
      </div>
    </div>
  )
}

export default FeedPage
import { useInfiniteQuery } from '@tanstack/react-query'
import { FeedsInfiniteQueryOptions } from '../../api/queries'
import type { PublicFeed } from '../../type'
import { FeedCardSkeleton } from '#/features/feeds/components/feedcard-skeleton'
import { useInfiniteScroll } from '#/hooks/useInfiiteScroll'
import FeedCard from '#/features/feeds/components/feedcard'
import PulsarLoader from '#/components/web/pulsar-loader'

export default function FeedsPage() {
  const query = useInfiniteQuery(FeedsInfiniteQueryOptions)

  const {
    items: feeds,
    bottomRef,
    isLoading,
    isFetchingNextPage,
  } = useInfiniteScroll<PublicFeed>(query)

  return (
    <div className="section-container max-w-2xl">
      <div className="space-y-4">
        <h3 className="text-xs text-muted-foreground font-semibold md:mt-3 mt-16">
          PUBLIC FEED
        </h3>
        <h1 className="md:text-4xl text-2xl font-bold text-foreground">
          What people are thinking.
        </h1>
      </div>

      <div className="my-10">
        <div className="flex flex-col gap-y-10">
          <FeedsList feeds={feeds} isLoading={isLoading} />
        </div>

        <div ref={bottomRef} />

        {isFetchingNextPage && (
          <PulsarLoader/>
        )}
      </div>
    </div>
  )
}

function FeedsList({
  feeds,
  isLoading,
}: {
  feeds: PublicFeed[]
  isLoading: boolean
}) {
  if (isLoading) {
    return <FeedCardSkeleton length={3} />
  }

  if (feeds.length === 0) {
    return <p className="text-muted-foreground md:text-sm text-xs">No feeds yet.</p>
  }

  return feeds.map((post) => <FeedCard key={post.id} {...post} />)
}

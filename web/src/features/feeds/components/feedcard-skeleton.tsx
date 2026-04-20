import { Skeleton } from '#/components/ui/skeleton'

export function FeedCardSkeleton({ length = 1 }: { length?: number }) {
  return (
    <>
      {Array.from({ length }).map((_, i) => (
        <div key={i} className="rounded-4xl bg-muted overflow-hidden">
          <div className="space-y-6 bg-card p-6">
            <Skeleton className="h-5 w-2/3" />
            <div className="space-y-2">
              <Skeleton className="h-3 w-full" />
              <Skeleton className="h-3 w-full" />
              <Skeleton className="h-3 w-4/5" />
            </div>
            <div className="flex gap-6">
              <Skeleton className="h-4 w-12" />
              <Skeleton className="h-4 w-12" />
            </div>
          </div>
          <div className="flex items-center justify-between p-6">
            <div className="flex items-center gap-3">
              <Skeleton className="w-8 h-8 rounded-full" />
              <Skeleton className="h-3 w-24" />
            </div>
            <Skeleton className="h-3 w-16" />
          </div>
        </div>
      ))}
    </>
  )
}

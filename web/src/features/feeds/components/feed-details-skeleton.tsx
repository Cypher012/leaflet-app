import { Skeleton } from '#/components/ui/skeleton'

export function FeedDetailSkeleton() {
  return (
    <div className="relative section-container max-w-3xl space-y-10 mt-12">
      {/* <div className="absolute top-14 left-0">
        <Skeleton className="size-10 rounded-full" />
      </div> */}

      <div className="p-5 space-y-7">
        <div className="flex items-center gap-2">
          <Skeleton className="size-2 rounded-full" />
          <Skeleton className="h-3 w-36" />
        </div>

        <div className="space-y-4">
          <div className="space-y-2.5">
            <Skeleton className="h-7 w-full" />
            <Skeleton className="h-7 w-2/3" />
          </div>
          <div className="flex items-center gap-5">
            <Skeleton className="size-10 rounded-full" />
            <Skeleton className="h-4 w-32" />
          </div>
        </div>

        <Skeleton className="w-full aspect-video rounded-2xl" />

        <div className="space-y-2">
          {Array.from({ length: 4 }).map((_, i) => (
            <Skeleton key={i} className="h-3 w-full" />
          ))}
          <Skeleton className="h-3 w-2/3" />
        </div>

        <div className="flex items-center gap-6">
          <div className="flex items-center gap-2">
            <Skeleton className="size-4 rounded" />
            <Skeleton className="h-3 w-7" />
          </div>
          <div className="flex items-center gap-2">
            <Skeleton className="size-4 rounded" />
            <Skeleton className="h-3 w-6" />
          </div>
        </div>
      </div>

      <div className="space-y-5 border-t pt-8">
        <Skeleton className="h-4 w-40" />
        <div className="flex gap-3">
          <Skeleton className="size-9 rounded-full shrink-0" />
          <Skeleton className="flex-1 h-20 rounded-xl" />
        </div>
      </div>
    </div>
  )
}

import { Skeleton } from '#/components/ui/skeleton'

const OverviewPageSkeleton = () => {
  return (
    <div className="flex w-full gap-3 max-w-5xl mx-auto py-10 sm:px-6 lg:px-4">
      {/* Main content */}
      <div className="flex-1 space-y-8">
        {/* Cards */}
        <div className="space-y-4">
          {Array.from({ length: 3 }).map((_, i) => (
            <div key={i} className="border border-border/10 rounded-2xl p-5 space-y-6">
              <div className="flex items-center gap-3">
                <Skeleton className="size-8 rounded-md" />
                <Skeleton className="h-3 w-48" />
              </div>
              <div className="border-l-2 border-border pl-4 space-y-2">
                <Skeleton className="h-4 w-full" />
                <Skeleton className="h-4 w-4/5" />
              </div>
              <div className="flex justify-between items-center">
                <Skeleton className="h-3 w-16" />
                <Skeleton className="h-3 w-12" />
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

export default OverviewPageSkeleton
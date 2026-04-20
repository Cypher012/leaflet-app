import { useEffect, useRef } from 'react'
import { useElementVisibility } from '@reactuses/core'
import type { UseInfiniteQueryResult } from '@tanstack/react-query'

type PaginatedResponse<T> = {
  data: T[]
}

export function useInfiniteScroll<T>(
  query: UseInfiniteQueryResult<{
    pages: PaginatedResponse<T>[]
  }>,
) {
  const { data, fetchNextPage, hasNextPage, isFetchingNextPage, isLoading } =
    query

  const items = data?.pages.flatMap((page) => page.data) ?? []

  const bottomRef = useRef<HTMLDivElement | null>(null)
  const [isVisible] = useElementVisibility(bottomRef)

  useEffect(() => {
    if (isVisible && hasNextPage && !isFetchingNextPage) {
      fetchNextPage()
    }
  }, [isVisible, hasNextPage, isFetchingNextPage, fetchNextPage])

  return {
    items,
    bottomRef,
    isLoading,
    isFetchingNextPage,
    hasNextPage,
  }
}

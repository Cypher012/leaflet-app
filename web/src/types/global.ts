export type PaginatedMeta = {
  count: number
  has_next: boolean
  next_cursor: string
}

export type PaginatedResponse<T> = {
  message: string
  data: T[]
  meta: PaginatedMeta
}

export type ApiResponse<T> = {
  message: string
  data: T
}

export type MessageResponse = {
  message: string
}

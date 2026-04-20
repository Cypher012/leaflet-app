// utils/time.ts

import { format } from 'timeago.js'

export function formatRelativeTime(date: string | Date): string {
  const d = typeof date === 'string' ? new Date(date) : date
  return format(d)
}

type Author = {
  username: string
  fullname: string
  avatar_url: string
}

export type PublicFeed = {
  id: string
  title: string
  content: string
  feed_image?: string
  is_liked: boolean
  author: Author
  stats: {
    likes: number
    comments: number
  }
  created_at: string
}

export type FeedComment = {
  id: string
  parent_id: string | null
  feed_id: string
  content: string
  author: Author
  stats: {
    like_count: number
    is_liked: boolean
  }
  created_at: string
  replies: FeedComment[]
}

export type CreateCommentReq = {
  content: string
}
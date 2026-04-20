export type ProfileUserResponse = {
    avatar_url: string,
    bio: string,
    fullname: string,
    id: string,
    stats: {
      comment_count: number,
      feed_count: number,
      like_count: number
    },
    username: string
}

export type ProfileOverviewResponse = {
      comment_body: string,
      content: string,
      created_at: string,
      id: string,
      parent_feed_title: string,
      feed_image: string,
      stats: {
        comment_count: number,
        is_liked: true,
        like_count: number
      },
      title: string,
      type: "feed" | "comment"
}

export type ProfileCommentResponse = {
    content: string,
    created_at: string,
    id: string,
    stats: {
    is_liked: boolean,
    like_count: number
    },
    title: string
}

export type ProfileFeedsResponse = {
    content: string,
    created_at: string,
    feed_image: string,
    id: string,
    stats: {
    comment_count: number,
    is_liked: true,
    like_count: number
    },
    title: string
}

export type FeedCardProps = {
  id: string,
  title: string,
  content: string,
  image: string,
  stats: {
    comment_count: number,
    is_liked: boolean,
    like_count: number
  },
  createdAt: string
}

export type CommentCardProps = {
  postTitle: string
  commentBody: string
  timestamp: string
  likes: number
}


export type CommentItemProps = {
  comment: Comment
  isReply?: boolean
  onReply: (parentId: number, content: string) => void
}


export type Comment = {
    id: string
    author: string
    avatar: string
    timeAgo: string
    content: string
    likes: number
    replies: Comment[]
  }
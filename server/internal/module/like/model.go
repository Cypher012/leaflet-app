package like

type FeedLikeExistsParams struct {
	UserID string
	FeedID string
}

type CommentLikeExistsParams struct {
	UserID    string
	CommentID string
}

type CreateFeedLikeParams struct {
	UserID string
	FeedID string
}

type CreateCommentLikeParams struct {
	UserID    string
	CommentID string
}

type DeleteFeedLikeParams struct {
	UserID string
	FeedID string
}

type DeleteCommentLikeParams struct {
	UserID    string
	CommentID string
}

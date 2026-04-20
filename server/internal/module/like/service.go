package like

import (
	"context"
)

type LikeService struct {
	repo  *LikeRepository
	limit int
}

func NewLikeService(repo *LikeRepository) *LikeService {
	return &LikeService{repo: repo, limit: 20}
}

func (s *LikeService) ToggleFeedLike(ctx context.Context, userID string, feedID string) error {
	isLiked, err := s.repo.FeedLikeExists(ctx, FeedLikeExistsParams{
		UserID: userID,
		FeedID: feedID,
	})
	if err != nil {
		return err
	}

	if !isLiked {
		return s.repo.CreateFeedLike(ctx, CreateFeedLikeParams{
			UserID: userID,
			FeedID: feedID,
		})
	}

	return s.repo.DeleteFeedLike(ctx, DeleteFeedLikeParams{
		UserID: userID,
		FeedID: feedID,
	})
}

func (s *LikeService) ToggleCommentLike(ctx context.Context, userID string, commentID string) error {
	isLiked, err := s.repo.CommentLikeExists(ctx, CommentLikeExistsParams{
		UserID:    userID,
		CommentID: commentID,
	})
	if err != nil {
		return err
	}

	if !isLiked {
		return s.repo.CreateCommentLike(ctx, CreateCommentLikeParams{
			UserID:    userID,
			CommentID: commentID,
		})
	}

	return s.repo.DeleteCommentLike(ctx, DeleteCommentLikeParams{
		UserID:    userID,
		CommentID: commentID,
	})
}

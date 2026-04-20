package like

import (
	"context"
	"server/internal/platform/db"
	"server/internal/shared/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LikeRepository struct {
	q    *db.Queries
	conn *pgxpool.Pool
}

func NewLikeRepository(conn *pgxpool.Pool) *LikeRepository {
	return &LikeRepository{
		q:    db.New(conn),
		conn: conn,
	}
}

func (r *LikeRepository) FeedLikeExists(ctx context.Context, params FeedLikeExistsParams) (bool, error) {
	feedUUID, err := utils.ConvertIdToPgUUID(&params.FeedID)
	if err != nil {
		return false, err
	}
	userUUID, err := utils.ConvertIdToPgUUID(&params.UserID)
	if err != nil {
		return false, err
	}

	return r.q.FeedLikeExists(ctx, db.FeedLikeExistsParams{
		UserID: userUUID,
		FeedID: feedUUID,
	})
}

func (r *LikeRepository) CommentLikeExists(ctx context.Context, params CommentLikeExistsParams) (bool, error) {
	commentUUID, err := utils.ConvertIdToPgUUID(&params.CommentID)
	if err != nil {
		return false, err
	}
	userUUID, err := utils.ConvertIdToPgUUID(&params.UserID)
	if err != nil {
		return false, err
	}

	return r.q.CommentLikeExists(ctx, db.CommentLikeExistsParams{
		UserID:    userUUID,
		CommentID: commentUUID,
	})
}

func (r *LikeRepository) CreateFeedLike(ctx context.Context, params CreateFeedLikeParams) error {
	feedUUID, err := utils.ConvertIdToPgUUID(&params.FeedID)
	if err != nil {
		return err
	}
	userUUID, err := utils.ConvertIdToPgUUID(&params.UserID)
	if err != nil {
		return err
	}

	return r.q.CreateFeedLike(ctx, db.CreateFeedLikeParams{
		UserID: userUUID,
		FeedID: feedUUID,
	})
}

func (r *LikeRepository) CreateCommentLike(ctx context.Context, params CreateCommentLikeParams) error {
	commentUUID, err := utils.ConvertIdToPgUUID(&params.CommentID)
	if err != nil {
		return err
	}
	userUUID, err := utils.ConvertIdToPgUUID(&params.UserID)
	if err != nil {
		return err
	}

	return r.q.CreateCommentLike(ctx, db.CreateCommentLikeParams{
		UserID:    userUUID,
		CommentID: commentUUID,
	})
}

func (r *LikeRepository) DeleteFeedLike(ctx context.Context, params DeleteFeedLikeParams) error {
	feedUUID, err := utils.ConvertIdToPgUUID(&params.FeedID)
	if err != nil {
		return err
	}
	userUUID, err := utils.ConvertIdToPgUUID(&params.UserID)
	if err != nil {
		return err
	}

	return r.q.DeleteFeedLike(ctx, db.DeleteFeedLikeParams{
		UserID: userUUID,
		FeedID: feedUUID,
	})
}

func (r *LikeRepository) DeleteCommentLike(ctx context.Context, params DeleteCommentLikeParams) error {
	commentUUID, err := utils.ConvertIdToPgUUID(&params.CommentID)
	if err != nil {
		return err
	}
	userUUID, err := utils.ConvertIdToPgUUID(&params.UserID)
	if err != nil {
		return err
	}

	return r.q.DeleteCommentLike(ctx, db.DeleteCommentLikeParams{
		UserID:    userUUID,
		CommentID: commentUUID,
	})
}

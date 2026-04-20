package comment

import (
	"context"
	"errors"
	"server/internal/platform/db"
	"server/internal/shared/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CommentRepository struct {
	q    *db.Queries
	conn *pgxpool.Pool
}

func NewCommentRepository(conn *pgxpool.Pool) *CommentRepository {
	return &CommentRepository{
		q:    db.New(conn),
		conn: conn,
	}
}

func (r *CommentRepository) CheckCommentExists(ctx context.Context, id string) (bool, error) {
	uuid, err := utils.ConvertIdToPgUUID(&id)
	if err != nil {
		return false, err
	}

	return r.q.CheckCommentExists(ctx, uuid)
}

func (r *CommentRepository) FeedComments(ctx context.Context, params FeedCommentsParams) ([]FeedComment, error) {
	viewerUUID, err := utils.ConvertIdToPgUUID(&params.ViewerID)
	if err != nil {
		return nil, err
	}

	feedUUID, err := utils.ConvertIdToPgUUID(&params.FeedID)
	if err != nil {
		return nil, err
	}

	cursorUUID, err := utils.ConvertIdToPgUUID(params.CursorID)
	if err != nil {
		return nil, err
	}

	dbComments, err := r.q.FeedComments(ctx, db.FeedCommentsParams{
		ViewerID:   viewerUUID,
		FeedID:     feedUUID,
		CursorID:   cursorUUID,
		CursorDate: utils.PgCursorTime(params.CursorDate),
		LimitCount: int32(params.LimitCount),
	})

	if err != nil {
		return nil, err
	}

	return NormalizeFeedComments(dbComments), err

}

func (r *CommentRepository) FeedCommentByID(ctx context.Context, params GetCommentByIDParams) (FeedComment, error) {
	commentUUID, err := utils.ConvertIdToPgUUID(&params.ID)
	if err != nil {
		return FeedComment{}, err
	}

	viewerUUID, err := utils.ConvertIdToPgUUID(&params.ViewerID)
	if err != nil {
		return FeedComment{}, err
	}

	dbComments, err := r.q.GetFeedCommentByID(ctx, db.GetFeedCommentByIDParams{
		ID:       commentUUID,
		ViewerID: viewerUUID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return FeedComment{}, ErrCommentNotFound
		}
		return FeedComment{}, err
	}

	return NormalizeFeedComment(dbComments), nil
}

func (r *CommentRepository) CreateComment(ctx context.Context, params CreateCommentParams) error {
	authorUUID, err := utils.ConvertIdToPgUUID(&params.AuthorID)
	if err != nil {
		return err
	}

	feedUUID, err := utils.ConvertIdToPgUUID(&params.FeedID)
	if err != nil {
		return err
	}

	parentUUID, err := utils.ConvertIdToPgUUID(&params.ParentID)
	if err != nil {
		return err
	}

	if err := r.q.CreateComment(ctx, db.CreateCommentParams{
		Content:  params.Content,
		AuthorID: authorUUID,
		FeedID:   feedUUID,
		ParentID: parentUUID,
	}); err != nil {
		return err
	}

	return nil
}

func (r *CommentRepository) CommentByID(ctx context.Context, id string) (Comment, error) {
	uuid, err := utils.ConvertIdToPgUUID(&id)
	if err != nil {
		return Comment{}, err
	}

	dbParentComments, err := r.q.GetCommentByID(ctx, uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Comment{}, ErrCommentNotFound
		}
		return Comment{}, err
	}

	return NormalizeComment(dbParentComments), nil
}

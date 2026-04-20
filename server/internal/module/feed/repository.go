package feed

import (
	"context"
	"errors"
	"server/internal/platform/db"
	"server/internal/shared/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FeedRepository struct {
	q    *db.Queries
	conn *pgxpool.Pool
}

func NewFeedRepository(conn *pgxpool.Pool) *FeedRepository {
	return &FeedRepository{
		q:    db.New(conn),
		conn: conn,
	}
}

func (r *FeedRepository) CheckFeedExists(ctx context.Context, feedID string) (bool, error) {
	uuid, err := utils.ConvertIdToPgUUID(&feedID)
	if err != nil {
		return false, err
	}
	return r.q.CheckFeedExists(ctx, uuid)
}

func (r *FeedRepository) GetFeeds(ctx context.Context, params GetFeedsParams) ([]Feed, error) {
	viewerUUID, err := utils.ConvertIdToPgUUID(params.ViewerID)
	if err != nil {
		return nil, err
	}

	cursorUUID, err := utils.ConvertIdToPgUUID(params.CursorID)
	if err != nil {
		return nil, err
	}

	dbFeeds, err := r.q.GetFeeds(ctx, db.GetFeedsParams{
		ViewerID:   viewerUUID,
		CursorDate: utils.PgCursorTime(params.CursorDate),
		CursorID:   cursorUUID,
		Limit:      int32(params.Limit),
	})
	if err != nil {
		return nil, err
	}

	return NormalizeFeeds(dbFeeds), nil
}

func (r *FeedRepository) FeedDetails(ctx context.Context, params FeedDetailsParams) (Feed, error) {
	feedUUID, err := utils.ConvertIdToPgUUID(&params.ID)
	if err != nil {
		return Feed{}, err
	}

	viewerUUID, err := utils.ConvertIdToPgUUID(params.ViewerID)
	if err != nil {
		return Feed{}, err
	}

	dbFeed, err := r.q.FeedDetails(ctx, db.FeedDetailsParams{
		ID:       feedUUID,
		ViewerID: viewerUUID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Feed{}, ErrFeedNotFound
		}
		return Feed{}, err
	}

	return NormalizeFeed(dbFeed), nil
}

func (r *FeedRepository) CreateFeed(ctx context.Context, params CreateFeedParams) error {
	userUUID, err := utils.ConvertIdToPgUUID(&params.UserID)
	if err != nil {
		return err
	}
	return r.q.CreateFeed(ctx, db.CreateFeedParams{
		UserID:    userUUID,
		Title:     params.Title,
		Content:   utils.PgText(params.Content),
		FeedImage: utils.PgText(params.ImageUrl),
	})
}

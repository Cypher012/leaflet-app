package profile

import (
	"context"
	"errors"
	"server/internal/platform/db"
	"server/internal/shared/types"
	"server/internal/shared/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepository struct {
	q    *db.Queries
	conn *pgxpool.Pool
}

func NewProfileRepository(conn *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{
		q:    db.New(conn),
		conn: conn,
	}
}

func (r *ProfileRepository) GetUserProfileByUsername(ctx context.Context, username string) (PublicUserProfile, error) {
	dbUser, err := r.q.GetUserProfileByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return PublicUserProfile{}, ErrUserNotFound
		}
		return PublicUserProfile{}, err
	}
	return NormalizePublicProfile(dbUser), nil
}

func (r *ProfileRepository) UpdateUserProfile(ctx context.Context, params db.UpdateUserProfileParams) (types.User, error) {
	dbUser, err := r.q.UpdateUserProfile(ctx, params)
	if err != nil {
		return types.User{}, err
	}
	return types.NormalizeUser(dbUser), nil
}

func (r *ProfileRepository) UserActivity(ctx context.Context, params UserActivityParams) ([]UserActivity, error) {
	userUUID, err := utils.ConvertIdToPgUUID(&params.UserID)
	if err != nil {
		return nil, err
	}

	viewerUUID, err := utils.ConvertIdToPgUUID(params.ViewerID)
	if err != nil {
		return nil, err
	}

	cursorUUID, err := utils.ConvertIdToPgUUID(params.CursorID)
	if err != nil {
		return nil, err
	}

	activity, err := r.q.UserActivity(ctx, db.UserActivityParams{
		UserID:     userUUID,
		ViewerID:   viewerUUID,
		CursorID:   cursorUUID,
		CursorDate: utils.PgCursorTime(params.CursorDate),
		Limit:      int32(params.Limit),
	})
	if err != nil {
		return nil, err
	}

	return NormalizeUserActivity(activity), nil
}

func (r *ProfileRepository) UserFeeds(ctx context.Context, params UserFeedsParams) ([]UserFeed, error) {
	userUUID, err := utils.ConvertIdToPgUUID(&params.UserID)
	if err != nil {
		return nil, err
	}

	viewerUUID, err := utils.ConvertIdToPgUUID(params.ViewerID)
	if err != nil {
		return nil, err
	}

	cursorUUID, err := utils.ConvertIdToPgUUID(params.CursorID)
	if err != nil {
		return nil, err
	}
	feeds, err := r.q.UserFeeds(ctx, db.UserFeedsParams{
		ViewerID:   viewerUUID,
		UserID:     userUUID,
		CursorID:   cursorUUID,
		CursorDate: utils.PgCursorTime(params.CursorDate),
		Limit:      int32(params.Limit),
	})
	if err != nil {
		return nil, err
	}

	return NormalizeUserFeeds(feeds), nil
}

func (r *ProfileRepository) UserComments(ctx context.Context, params UserCommentsParams) ([]UserComment, error) {
	userUUID, err := utils.ConvertIdToPgUUID(&params.UserID)
	if err != nil {
		return nil, err
	}

	viewerUUID, err := utils.ConvertIdToPgUUID(params.ViewerID)
	if err != nil {
		return nil, err
	}

	cursorUUID, err := utils.ConvertIdToPgUUID(params.CursorID)
	if err != nil {
		return nil, err
	}
	comments, err := r.q.UserComments(ctx, db.UserCommentsParams{
		ViewerID:   viewerUUID,
		UserID:     userUUID,
		CursorID:   cursorUUID,
		CursorDate: utils.PgCursorTime(params.CursorDate),
		Limit:      int32(params.Limit),
	})

	if err != nil {
		return nil, err
	}

	return NormalizeUserComments(comments), nil
}

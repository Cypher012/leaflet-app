package profile

import (
	"context"
	"log"
	"server/internal/shared/response"
	"server/internal/shared/utils"
	"time"
)

type ProfileService struct {
	repo  *ProfileRepository
	limit int
}

func NewProfileService(repo *ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo, limit: 10}
}

func (s *ProfileService) GetUserFromUsername(ctx context.Context, username string) (PublicUserProfile, error) {
	user, err := s.repo.GetUserProfileByUsername(ctx, username)
	if err != nil {
		return PublicUserProfile{}, err
	}
	return user, nil
}

func (s *ProfileService) GetUserActivity(ctx context.Context, userID string, viewerID, cursor string) ([]UserActivity, response.PaginationMeta, error) {
	cursorDate, cursorId, err := utils.ParseCursor(cursor)
	if err != nil {
		return nil, response.PaginationMeta{}, err
	}

	log.Println("viewerID :", &viewerID)

	activity, err := s.repo.UserActivity(ctx, UserActivityParams{
		UserID:     userID,
		ViewerID:   &viewerID,
		CursorID:   cursorId,
		CursorDate: cursorDate,
		Limit:      s.limit + 1,
	})

	activity, meta := utils.BuildNextCursor(activity, s.limit, func(r UserActivity) (string, time.Time) {
		return r.ID, r.CreatedAt
	})

	return activity, meta, err
}

func (s *ProfileService) GetUserFeeds(ctx context.Context, userID string, viewerID, cursor string) ([]UserFeed, response.PaginationMeta, error) {
	cursorDate, cursorId, err := utils.ParseCursor(cursor)
	if err != nil {
		return nil, response.PaginationMeta{}, err
	}

	feeds, err := s.repo.UserFeeds(ctx, UserFeedsParams{
		UserID:     userID,
		ViewerID:   &viewerID,
		CursorID:   cursorId,
		CursorDate: cursorDate,
		Limit:      s.limit + 1,
	})

	feeds, meta := utils.BuildNextCursor(feeds, s.limit, func(r UserFeed) (string, time.Time) {
		return r.ID, r.CreatedAt
	})

	return feeds, meta, err
}

func (s *ProfileService) GetUserComments(ctx context.Context, userID string, viewerID, cursor string) ([]UserComment, response.PaginationMeta, error) {
	cursorDate, cursorId, err := utils.ParseCursor(cursor)
	if err != nil {
		return nil, response.PaginationMeta{}, err
	}

	comments, err := s.repo.UserComments(ctx, UserCommentsParams{
		UserID:     userID,
		ViewerID:   &viewerID,
		CursorID:   cursorId,
		CursorDate: cursorDate,
		Limit:      s.limit + 1,
	})

	comments, meta := utils.BuildNextCursor(comments, s.limit, func(r UserComment) (string, time.Time) {
		return r.ID, r.CreatedAt
	})

	return comments, meta, err
}

package feed

import (
	"context"
	"server/internal/shared/response"
	"server/internal/shared/utils"
	"time"
)

type FeedService struct {
	repo  *FeedRepository
	limit int
}

func NewFeedService(repo *FeedRepository) *FeedService {
	return &FeedService{repo: repo, limit: 20}
}

func (s *FeedService) CheckIfFeedExist(ctx context.Context, feed_id string) error {
	isFeed, err := s.repo.CheckFeedExists(ctx, feed_id)
	if err != nil {
		return err
	}
	if !isFeed {
		return ErrFeedNotFound
	}
	return nil
}

func (s *FeedService) FetchFeeds(ctx context.Context, cursor string, viewerID string) (feeds []Feed, meta response.PaginationMeta, err error) {
	cursorDate, cursorId, err := utils.ParseCursor(cursor)

	if err != nil {
		return nil, response.PaginationMeta{}, err
	}

	feeds, err = s.repo.GetFeeds(ctx, GetFeedsParams{
		ViewerID:   &viewerID,
		CursorDate: cursorDate,
		CursorID:   cursorId,
		Limit:      s.limit + 1,
	})

	if err != nil {
		return nil, response.PaginationMeta{}, err
	}

	feeds, meta = utils.BuildNextCursor(feeds, s.limit, func(r Feed) (string, time.Time) {
		return r.ID, r.CreatedAt
	})

	return feeds, meta, nil
}

func (s *FeedService) FetchFeed(ctx context.Context, feedId string, viewerID string) (Feed, error) {
	feed, err := s.repo.FeedDetails(ctx, FeedDetailsParams{
		ID:       feedId,
		ViewerID: &viewerID,
	})

	if err != nil {
		return Feed{}, err
	}

	return feed, nil
}

func (s *FeedService) CreateFeed(ctx context.Context, userID string, body CreateFeedReq) error {
	err := s.repo.CreateFeed(ctx, CreateFeedParams{
		UserID:   userID,
		Title:    body.Title,
		Content:  body.Content,
		ImageUrl: body.FeedImage,
	})

	if err != nil {
		return err
	}

	return nil
}

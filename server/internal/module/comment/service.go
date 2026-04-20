package comment

import (
	"context"
	"server/internal/shared/response"
	"server/internal/shared/utils"
	"time"
)

type CommentService struct {
	repo  *CommentRepository
	limit int
}

func NewCommentService(repo *CommentRepository) *CommentService {
	return &CommentService{repo: repo, limit: 20}
}

func (s *CommentService) CheckIfCommentExist(ctx context.Context, comment_id string) error {
	isComment, err := s.repo.CheckCommentExists(ctx, comment_id)
	if err != nil {
		return err
	}
	if !isComment {
		return ErrCommentNotFound
	}
	return nil
}

func (s *CommentService) FetchComments(
	ctx context.Context,
	viewerID string,
	feedID string,
	cursor string,
) (comments []*FeedComment, meta response.PaginationMeta, err error) {

	cursorDate, cursorId, err := utils.ParseCursor(cursor)
	if err != nil {
		return nil, response.PaginationMeta{}, err
	}

	feedComments, err := s.repo.FeedComments(ctx, FeedCommentsParams{
		ViewerID:   viewerID,
		FeedID:     feedID,
		CursorDate: cursorDate,
		CursorID:   cursorId,
		LimitCount: s.limit + 1,
	})

	if err != nil {
		return nil, response.PaginationMeta{}, err
	}

	feedComments, meta = utils.BuildNextCursor(feedComments, s.limit, func(r FeedComment) (string, time.Time) {
		return r.ID, r.CreatedAt
	})

	return NormalizeNestComments(feedComments), meta, nil
}

func (s *CommentService) PostFeedComment(ctx context.Context, userID string, feedId string, content string) error {
	if err := s.repo.CreateComment(ctx, CreateCommentParams{
		Content:  content,
		AuthorID: userID,
		FeedID:   feedId,
		ParentID: "",
	}); err != nil {
		return err
	}

	return nil
}

func (s *CommentService) ReplyComment(ctx context.Context, authorID string, commentID string, content string) error {
	parent, err := s.repo.CommentByID(ctx, commentID)
	if err != nil {
		return err
	}

	if err := s.repo.CreateComment(ctx, CreateCommentParams{
		Content:  content,
		AuthorID: authorID,
		FeedID:   parent.FeedID,
		ParentID: commentID,
	}); err != nil {
		return err
	}

	return nil
}

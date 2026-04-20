package comment

import (
	"errors"
	"server/internal/platform/db"
	"server/internal/shared/response"
	"time"
)

// Error type
var ErrCommentNotFound = errors.New("comment not found")

// JSON Body request
type NewComment struct {
	Content string `json:"content" validator:"required,min=1"`
}

// Swagger type
type DocCommentsResponse response.PaginatedResponse[*FeedComment]

// DB Params
type FeedCommentsParams struct {
	ViewerID   string
	FeedID     string
	CursorDate *time.Time
	CursorID   *string
	LimitCount int
}

type GetCommentByIDParams struct {
	ViewerID string
	ID       string
}

type CreateCommentParams struct {
	Content  string
	AuthorID string
	FeedID   string
	ParentID string
}

// DB Model
type Comment struct {
	ID        string
	AuthorID  string
	FeedID    string
	ParentID  *string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FeedCommentAuthor struct {
	Fullname  string `json:"fullname"`
	AvatarUrl string `json:"avatar_url"`
	Username  string `json:"username"`
}

type FeedCommentStats struct {
	LikeCount int64 `json:"like_count"`
	IsLiked   bool  `json:"is_liked"`
}

type FeedComment struct {
	ID        string            `json:"id"`
	ParentID  *string           `json:"parent_id"`
	FeedID    string            `json:"feed_id"`
	CreatedAt time.Time         `json:"created_at"`
	Content   string            `json:"content"`
	Author    FeedCommentAuthor `json:"author"`
	Stats     FeedCommentStats  `json:"stats"`
	Replies   []*FeedComment    `json:"replies"`
}

func NormalizeComment(c db.Comment) Comment {
	var parentID *string
	if c.ParentID.Valid {
		s := c.ParentID.String()
		parentID = &s
	}

	return Comment{
		ID:        c.ID.String(),
		ParentID:  parentID,
		FeedID:    c.FeedID.String(),
		CreatedAt: c.CreatedAt.Time,
		Content:   c.Content,
		AuthorID:  c.AuthorID.String(),
		UpdatedAt: c.UpdatedAt.Time,
	}
}

func NormalizeFeedComment(c db.GetFeedCommentByIDRow) FeedComment {
	var parentID *string
	if c.ParentID.Valid {
		s := c.ParentID.String()
		parentID = &s
	}

	return FeedComment{
		ID:        c.ID.String(),
		ParentID:  parentID,
		FeedID:    c.FeedID.String(),
		CreatedAt: c.CreatedAt.Time,
		Content:   c.Content,
		Author: FeedCommentAuthor{
			Fullname:  c.Fullname.String,
			Username:  c.Username.String,
			AvatarUrl: c.AvatarUrl.String,
		},
		Stats: FeedCommentStats{
			LikeCount: c.LikeCount,
			IsLiked:   c.IsLiked,
		},
	}
}

func NormalizeFeedComments(comments []db.FeedCommentsRow) []FeedComment {
	result := make([]FeedComment, 0, len(comments))

	for _, c := range comments {
		var parentID *string
		if c.ParentID.Valid {
			s := c.ParentID.String()
			parentID = &s
		}

		result = append(result, FeedComment{
			ID:        c.ID.String(),
			ParentID:  parentID,
			FeedID:    c.FeedID.String(),
			CreatedAt: c.CreatedAt.Time,
			Content:   c.Content,
			Author: FeedCommentAuthor{
				Fullname:  c.Fullname.String,
				Username:  c.Username.String,
				AvatarUrl: c.AvatarUrl.String,
			},
			Stats: FeedCommentStats{
				LikeCount: c.LikeCount,
				IsLiked:   c.IsLiked,
			},
		})
	}
	return result
}

func NormalizeNestComments(comments []FeedComment) []*FeedComment {
	commentMap := make(map[string]*FeedComment, 0)

	for _, c := range comments {
		commentMap[c.ID] = &FeedComment{
			ID:        c.ID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
			ParentID:  c.ParentID,
			FeedID:    c.FeedID,
			Author: FeedCommentAuthor{
				Fullname:  c.Author.Fullname,
				Username:  c.Author.Username,
				AvatarUrl: c.Author.AvatarUrl,
			},
			Stats: FeedCommentStats{
				LikeCount: c.Stats.LikeCount,
				IsLiked:   c.Stats.IsLiked,
			},
			Replies: []*FeedComment{},
		}
	}

	topLevel := make([]*FeedComment, 0)
	for _, row := range comments {
		c := commentMap[row.ID]
		if row.ParentID == nil {
			topLevel = append(topLevel, c)
			continue
		}

		if parent, ok := commentMap[*row.ParentID]; ok {
			parent.Replies = append(parent.Replies, c)
		}
	}

	return topLevel
}

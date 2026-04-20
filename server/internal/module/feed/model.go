package feed

import (
	"errors"
	"server/internal/platform/db"
	"server/internal/shared/response"
	"time"
)

// Handler request
type CreateFeedReq struct {
	Title     string `json:"title" validate:"required,min=1,max=120"`
	Content   string `json:"content" validate:"required,min=1,max=5000"`
	FeedImage string `json:"feed_image" validate:"omitempty,url"`
}

// Error types
var ErrFeedNotFound = errors.New("feed not found")

// DB Params

type GetFeedsParams struct {
	ViewerID   *string
	CursorDate *time.Time
	CursorID   *string
	Limit      int
}

type FeedDetailsParams struct {
	ID       string
	ViewerID *string
}

type CreateFeedParams struct {
	UserID   string
	Title    string
	Content  string
	ImageUrl string
}

type GetUserFeedsParams struct {
	UserID     string
	ViewerID   *string
	CursorDate *time.Time
	CursorID   *string
	Limit      int
}

// Swagger Response
type DocFeedResponse response.Response[Feed]
type DocFeedsResponse response.PaginatedResponse[Feed]

type FeedAuthor struct {
	Username  string `json:"username"`
	Fullname  string `json:"fullname"`
	AvatarUrl string `json:"avatar_url"`
}

type FeedStats struct {
	LikesCount   int `json:"likes"`
	CommentCount int `json:"comments"`
}

type Feed struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	FeedImage string     `json:"feed_image"`
	Author    FeedAuthor `json:"author"`
	Stats     FeedStats  `json:"stats"`
	IsLiked   bool       `json:"is_liked"`
	CreatedAt time.Time  `json:"created_at"`
}

func NormalizeFeed(feed db.FeedDetailsRow) Feed {
	return Feed{
		ID:        feed.ID.String(),
		Title:     feed.Title,
		Content:   feed.Content.String,
		FeedImage: feed.FeedImage.String,
		CreatedAt: feed.CreatedAt.Time,
		IsLiked:   feed.IsLiked,
		Author: FeedAuthor{
			Fullname:  feed.Fullname.String,
			Username:  feed.Username.String,
			AvatarUrl: feed.AvatarUrl.String,
		},
		Stats: FeedStats{
			LikesCount:   int(feed.LikeCount),
			CommentCount: int(feed.CommentCount),
		},
	}
}

func NormalizeFeeds(feeds []db.GetFeedsRow) []Feed {
	result := make([]Feed, 0, len(feeds))

	for _, f := range feeds {
		result = append(result, Feed{
			ID:        f.ID.String(),
			Title:     f.Title,
			Content:   f.Content.String,
			FeedImage: f.FeedImage.String,
			CreatedAt: f.CreatedAt.Time,
			IsLiked:   f.IsLiked,
			Author: FeedAuthor{
				Fullname:  f.Fullname.String,
				Username:  f.Username.String,
				AvatarUrl: f.AvatarUrl.String,
			},
			Stats: FeedStats{
				LikesCount:   int(f.LikeCount),
				CommentCount: int(f.CommentCount),
			},
		})
	}
	return result
}

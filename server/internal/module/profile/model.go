package profile

import (
	"errors"
	"server/internal/platform/db"
	"server/internal/shared/response"
	"time"
)

type DocUserProfileResponse response.Response[PublicUserProfile]
type DocUserActivityResponse response.PaginatedResponse[UserActivity]
type DocUserPublicFeedResponse response.PaginatedResponse[UserFeed]
type DocUserFeedCommentResponse response.PaginatedResponse[UserComment]

var ErrUserNotFound = errors.New("user not found")

type UpdateUserProfileParams struct {
	Fullname string
	Username string
	Bio      *string
	ID       string
}

type UserActivityParams struct {
	UserID     string
	ViewerID   *string
	CursorID   *string
	CursorDate *time.Time
	Limit      int
}

type UserFeedsParams struct {
	UserID     string
	ViewerID   *string
	CursorID   *string
	CursorDate *time.Time
	Limit      int
}

type UserCommentsParams struct {
	UserID     string
	ViewerID   *string
	CursorID   *string
	CursorDate *time.Time
	Limit      int
}

type UserStats struct {
	LikeCount    int  `json:"like_count"`
	CommentCount int  `json:"comment_count"`
	IsLiked      bool `json:"is_liked"`
}

type ProfileUserStats struct {
	LikeCount    int `json:"like_count"`
	CommentCount int `json:"comment_count"`
	FeedCount    int `json:"feed_count"`
}

type ProfileUserCommentStats struct {
	IsLiked   bool `json:"is_liked"`
	LikeCount int  `json:"like_count"`
}

type UserActivity struct {
	Type            string    `json:"type"`
	ID              string    `json:"id"`
	Title           string    `json:"title,omitempty"`
	Content         string    `json:"content,omitempty"`
	FeedImage       string    `json:"feed_image,omitempty"`
	CommentBody     string    `json:"comment_body,omitempty"`
	ParentFeedTitle string    `json:"parent_feed_title,omitempty"`
	Stats           UserStats `json:"stats"`
	CreatedAt       time.Time `json:"created_at"`
}

type UserFeed struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	FeedImage string    `json:"feed_image"`
	Stats     UserStats `json:"stats"`
	CreatedAt time.Time `json:"created_at"`
}

type UserComment struct {
	ID        string                  `json:"id"`
	Title     string                  `json:"title,omitempty"`
	Content   string                  `json:"content"`
	CreatedAt time.Time               `json:"created_at"`
	Stats     ProfileUserCommentStats `json:"stats"`
}

type PublicUserProfile struct {
	ID        string           `json:"id"`
	Fullname  string           `json:"fullname"`
	Username  string           `json:"username"`
	Bio       string           `json:"bio,omitempty"`
	AvatarUrl string           `json:"avatar_url"`
	Stats     ProfileUserStats `json:"stats"`
}

func NormalizePublicProfile(u db.GetUserProfileByUsernameRow) PublicUserProfile {
	return PublicUserProfile{
		ID:        u.ID.String(),
		Fullname:  u.Fullname,
		Username:  u.Username,
		Bio:       u.Bio.String,
		AvatarUrl: u.AvatarUrl.String,
		Stats: ProfileUserStats{
			LikeCount:    int(u.LikeCount),
			CommentCount: int(u.CommentCount),
			FeedCount:    int(u.FeedCount),
		},
	}
}

func NormalizeUserActivity(rows []db.UserActivityRow) []UserActivity {
	activities := make([]UserActivity, 0, len(rows))

	for _, a := range rows {
		activities = append(activities, UserActivity{
			Type:            a.Type,
			ID:              a.ID.String(),
			Title:           a.Title,
			Content:         a.Content.String,
			CommentBody:     a.CommentBody,
			ParentFeedTitle: a.ParentFeedTitle,
			FeedImage:       a.FeedImage.String,
			CreatedAt:       a.CreatedAt.Time,
			Stats: UserStats{
				LikeCount:    int(a.LikeCount),
				CommentCount: int(a.CommentCount),
				IsLiked:      a.IsLiked,
			},
		})
	}

	return activities
}

func NormalizeUserFeeds(rows []db.UserFeedsRow) []UserFeed {
	feeds := make([]UserFeed, 0, len(rows))

	for _, f := range rows {
		feeds = append(feeds, UserFeed{
			ID:        f.ID.String(),
			Title:     f.Title,
			Content:   f.Content.String,
			FeedImage: f.FeedImage.String,
			CreatedAt: f.CreatedAt.Time,
			Stats: UserStats{
				LikeCount:    int(f.LikeCount),
				IsLiked:      f.IsLiked,
				CommentCount: int(f.CommentCount),
			},
		})
	}

	return feeds
}

func NormalizeUserComments(rows []db.UserCommentsRow) []UserComment {
	comments := make([]UserComment, 0, len(rows))

	for _, c := range rows {
		comments = append(comments, UserComment{
			ID:        c.ID.String(),
			Title:     c.Title.String,
			Content:   c.Content,
			CreatedAt: c.CreatedAt.Time,
			Stats: ProfileUserCommentStats{
				LikeCount: int(c.LikeCount),
				IsLiked:   c.IsLiked,
			},
		})
	}

	return comments
}

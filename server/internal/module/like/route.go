package like

import (
	"server/internal/module/comment"
	"server/internal/module/feed"
	"server/internal/shared/middleware"
	"server/internal/shared/types"

	"github.com/labstack/echo/v5"
)

func LikeRoute(r *echo.Group, deps types.RouterDeps, authStore middleware.AuthStore) *echo.Group {
	feedRepository := feed.NewFeedRepository(deps.Conn)
	feedService := feed.NewFeedService(feedRepository)

	commentRepository := comment.NewCommentRepository(deps.Conn)
	commentService := comment.NewCommentService(commentRepository)

	repo := NewLikeRepository(deps.Conn)
	service := NewLikeService(repo)
	handler := NewLikeHandler(service, feedService, commentService, deps.Logger, deps.Config)

	protected := r.Group("")
	protected.Use(middleware.RequireAuth(authStore, deps.Logger))

	// Like routes
	protected.POST("/feeds/:feed_id/like", handler.ToggleFeedLike)
	protected.POST("/comments/:comment_id/like", handler.ToggleCommentLike)

	return r
}

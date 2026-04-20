package comment

import (
	"server/internal/module/feed"
	"server/internal/shared/middleware"
	"server/internal/shared/types"

	"github.com/labstack/echo/v5"
)

func CommentRoute(r *echo.Group, deps types.RouterDeps, authStore middleware.AuthStore) *echo.Group {
	feedRepository := feed.NewFeedRepository(deps.Conn)
	feedService := feed.NewFeedService(feedRepository)

	repo := NewCommentRepository(deps.Conn)
	service := NewCommentService(repo)
	handler := NewCommentHandler(service, feedService, deps.Logger, deps.Config)

	r.Use(middleware.OptionalAuth(authStore, deps.Logger))

	r.GET("/feeds/:feed_id/comments", handler.GetFeedsComments)

	protected := r.Group("")
	protected.Use(middleware.RequireAuth(authStore, deps.Logger))

	protected.POST("/feeds/:feed_id/comments", handler.PostComment)
	protected.POST("/comments/:comment_id/replies", handler.ReplyComment)

	return r
}

// ideally what should be the maximum timeout an api requewst shoudl have

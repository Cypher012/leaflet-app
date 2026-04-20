package feed

import (
	"server/internal/shared/middleware"
	"server/internal/shared/types"

	"github.com/labstack/echo/v5"
)

func FeedRoute(r *echo.Group, deps types.RouterDeps, authStore middleware.AuthStore) *echo.Group {

	repo := NewFeedRepository(deps.Conn)
	service := NewFeedService(repo)
	handler := NewFeedHandler(service, deps.Logger, deps.Config)

	r.Use(middleware.OptionalAuth(authStore, deps.Logger))

	r.GET("/feeds", handler.PublicFeeds)
	r.GET("/feeds/:id", handler.PublicFeed)

	protected := r.Group("")
	protected.Use(middleware.RequireAuth(authStore, deps.Logger))

	protected.POST("/feeds", handler.CreateFeed)

	return r
}

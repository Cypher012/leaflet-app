package profile

import (
	"server/internal/module/auth"
	"server/internal/shared/middleware"
	"server/internal/shared/types"

	"github.com/labstack/echo/v5"
)

func ProfileRoute(r *echo.Group, deps types.RouterDeps, authRepo *auth.AuthRepository, authStore middleware.AuthStore) *echo.Group {
	authService := auth.NewAuthService(authRepo)

	repo := NewProfileRepository(deps.Conn)
	service := NewProfileService(repo)
	handler := NewProfileHandler(service, authService, deps.Logger, deps.Config)

	r.Use(middleware.OptionalAuth(authStore, deps.Logger))

	r.GET("/profile/:username", handler.UserProfile)
	r.GET("/profile/:username/overview", handler.UserProfileOverview)
	r.GET("/profile/:username/feeds", handler.UserProfileFeeds)
	r.GET("/profile/:username/comments", handler.UserProfileComments)

	return r
}

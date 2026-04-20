package main

import (
	"net/http"
	"server/internal/module/auth"
	"server/internal/module/comment"
	"server/internal/module/feed"
	"server/internal/module/like"
	"server/internal/module/profile"
	"server/internal/module/upload"
	"server/internal/shared/middleware"
	"server/internal/shared/types"
	"server/internal/shared/utils"
	"server/views"

	_ "server/docs"

	"github.com/labstack/echo/v5"
	echoSwagger "github.com/swaggo/echo-swagger/v2"
)

func NewRouter(r *echo.Echo, deps types.RouterDeps) {
	r.GET("/", func(c *echo.Context) error {
		url := deps.Config.BackendURL
		return utils.Render(c, http.StatusOK, views.HomePage(url))
	})

	r.GET("/docs/*", echoSwagger.WrapHandler)

	api := r.Group("api")

	authRepo := auth.NewAuthRepository(deps.Conn)
	var authStore middleware.AuthStore = authRepo

	// Auth Route
	auth.AuthRoute(api, deps, authRepo)

	// Feed Route
	feed.FeedRoute(api, deps, authStore)

	// Comment Route
	comment.CommentRoute(api, deps, authStore)

	// Like Route
	like.LikeRoute(api, deps, authStore)

	// Profile Route
	profile.ProfileRoute(api, deps, authRepo, authStore)

	// Upload Route
	upload.UploadRoute(api, deps, authStore)

}

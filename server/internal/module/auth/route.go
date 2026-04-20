package auth

import (
	"server/internal/shared/types"

	"github.com/labstack/echo/v5"
)

func AuthRoute(r *echo.Group, deps types.RouterDeps, authRepo *AuthRepository) *echo.Group {

	service := NewAuthService(authRepo)
	handler := NewAuthHandler(service, deps.Logger, deps.Config)

	r.GET("/auth/:provider", handler.OAuthLogin)
	r.GET("/auth/:provider/callback", handler.OAuthCallback)
	r.GET("/auth/session", handler.GetCurrentSession)
	r.GET("/auth/me", handler.GetMe)
	r.POST("/auth/logout", handler.Logout)

	return r
}

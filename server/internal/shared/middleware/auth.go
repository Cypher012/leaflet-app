package middleware

import (
	"context"
	"errors"
	"log/slog"
	"server/internal/module/auth"
	"server/internal/shared/response"
	"server/internal/shared/types"

	"github.com/labstack/echo/v5"
)

type AuthStore interface {
	GetSessionByToken(ctx context.Context, token string) (types.Session, error)
	GetUserByID(ctx context.Context, id string) (types.User, error)
}

func OptionalAuth(store AuthStore, logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			token, err := c.Cookie("leaflet_sid")
			if err != nil {
				// no cookie → just continue (guest)
				return next(c)
			}

			ctx := c.Request().Context()

			session, err := store.GetSessionByToken(ctx, token.Value)
			if err != nil {
				// invalid session → treat as guest
				return next(c)
			}

			user, err := store.GetUserByID(ctx, session.UserID)
			if err != nil {
				// optional: log but don't fail request
				logger.Error("failed to fetch user", slog.Any("error", err))
				return next(c)
			}

			c.Set("user", &user)

			return next(c)
		}
	}
}

func RequireAuth(store AuthStore, logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			token, err := c.Cookie("leaflet_sid")
			if err != nil {
				return response.UnauthorizedError(c, logger, nil)
			}

			ctx := c.Request().Context()

			session, err := store.GetSessionByToken(ctx, token.Value)
			if err != nil {
				if errors.Is(err, auth.ErrSessionNotFound) {
					return response.UnauthorizedError(c, logger, nil)
				}
				return response.InternalError(c, logger, err)
			}

			user, err := store.GetUserByID(ctx, session.UserID)
			if err != nil {
				return response.InternalError(c, logger, err)
			}

			c.Set("user", &user)
			return next(c)
		}
	}
}

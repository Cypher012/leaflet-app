package middleware

import (
	"log/slog"

	"github.com/labstack/echo/v5"
)

func LoggerMiddleware(log *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			req := c.Request()

			reqLogger := log.With(
				slog.Group("request",
					slog.String("method", req.Method),
					slog.String("path", req.URL.Path),
				),
			)

			c.Set("logger", reqLogger)

			reqLogger.Info("request started")

			err := next(c)

			if err != nil {
				reqLogger.Error("request failed", "error", err)
			} else {
				reqLogger.Info("Request completed")
			}

			return err
		}
	}
}

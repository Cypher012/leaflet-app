package response

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
)

func OK[T any](c *echo.Context, message string, data T) error {
	return c.JSON(http.StatusOK, NewResponse(message, data))
}

func Created(c *echo.Context, message string) error {
	return c.JSON(http.StatusCreated, NewMessageResponse(message))
}

func CreatedWithData[T any](c *echo.Context, message string, data T) error {
	return c.JSON(http.StatusCreated, NewResponse(message, data))
}

func OKMessage(c *echo.Context, message string) error {
	return c.JSON(http.StatusOK, Response[any]{
		Message: message,
	})
}

func Redirect(c *echo.Context, url string) error {
	return c.Redirect(http.StatusFound, url)
}

func Paginated[T any](c *echo.Context, message string, data []T, meta PaginationMeta) error {
	return c.JSON(http.StatusOK, NewPaginatedResponse(message, data, meta))
}

func Fail(c *echo.Context, log *slog.Logger, status int, message string, err error) error {
	log.Error("request failed",
		slog.Int("status", status),
		slog.String("message", message),
		slog.String("path", c.Request().URL.Path),
		slog.String("method", c.Request().Method),
		slog.Any("error", err),
	)
	return c.JSON(status, NewErrorResponse(message))
}

func BadRequest(c *echo.Context, log *slog.Logger, message string, err error) error {
	return Fail(c, log, http.StatusBadRequest, message, err)
}

func NotFound(c *echo.Context, log *slog.Logger, message string, err error) error {
	return Fail(c, log, http.StatusNotFound, message, err)
}

func UnauthorizedError(c *echo.Context, log *slog.Logger, err error) error {
	return Fail(c, log, http.StatusUnauthorized, "unauthorized request", err)
}

func InternalError(c *echo.Context, log *slog.Logger, err error) error {
	return Fail(c, log, http.StatusInternalServerError, "internal error: something went wrong", err)
}

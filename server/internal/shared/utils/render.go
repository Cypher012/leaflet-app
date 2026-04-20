package utils

import (
	"context"
	"io"

	"github.com/labstack/echo/v5"
)

type TemplComponent interface {
	Render(ctx context.Context, w io.Writer) error
}

func Render(c *echo.Context, status int, comp TemplComponent) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	c.Response().WriteHeader(status)
	return comp.Render(c.Request().Context(), c.Response())
}

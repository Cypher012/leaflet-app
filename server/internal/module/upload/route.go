package upload

import (
	"server/internal/shared/middleware"
	"server/internal/shared/types"

	"github.com/labstack/echo/v5"
)

func UploadRoute(r *echo.Group, deps types.RouterDeps, authStore middleware.AuthStore) *echo.Group {

	service := NewUploadService(deps.S3Storage)
	handler := NewUploadHandler(service, deps.Logger)

	upload := r.Group("/upload", middleware.RequireAuth(authStore, deps.Logger))
	upload.GET("/presign", handler.GetPresignedURL)

	return r
}

package upload

import (
	"log/slog"
	"server/internal/shared/response"

	"github.com/labstack/echo/v5"
)

type UploadHandler struct {
	service *UploadService
	logger  *slog.Logger
}

func NewUploadHandler(service *UploadService, logger *slog.Logger) *UploadHandler {
	return &UploadHandler{service, logger}
}

// GetPresignedURL godoc
// @Summary      Get presigned upload URL
// @Description  Returns a presigned URL for direct upload to R2 storage and the public URL to save in the database
// @Tags         upload
// @Accept       json
// @Produce      json
// @Param        type          query     string  true  "Upload type (avatar, feed)"
// @Param        content_type  query     string  true  "File content type (e.g. image/png, image/jpeg)"
// @Success      200           {object}  DocPresignResponse
// @Failure      400           {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /upload/presign [get]
func (h *UploadHandler) GetPresignedURL(c *echo.Context) error {
	uploadType := c.QueryParam("type")
	contentType := c.QueryParam("content_type")

	if uploadType == "" || contentType == "" {
		return response.BadRequest(c, h.logger, "type and content_type are required", nil)
	}

	uploadURL, publicURL, err := h.service.GetPresignedURL(c.Request().Context(), uploadType, contentType)
	if err != nil {
		return response.BadRequest(c, h.logger, err.Error(), err)
	}

	return response.OK(c, "presigned url fetched successfully", PresignResponse{
		UploadURL: uploadURL,
		PublicURL: publicURL,
	})
}

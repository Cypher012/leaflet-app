package like

import (
	"errors"
	"log/slog"
	"server/internal/module/comment"
	"server/internal/module/feed"
	"server/internal/platform/config"
	"server/internal/shared/response"
	"server/internal/shared/utils"

	"github.com/labstack/echo/v5"
)

type LikeHandler struct {
	service        *LikeService
	feedService    *feed.FeedService
	commentService *comment.CommentService
	logger         *slog.Logger
	cfg            config.AppConfig
}

func NewLikeHandler(service *LikeService, feedService *feed.FeedService, commentService *comment.CommentService, logger *slog.Logger, cfg config.AppConfig) *LikeHandler {
	return &LikeHandler{service, feedService, commentService, logger, cfg}
}

// ToggleFeedLike godoc
// @Summary      Toggle feed like
// @Description  Like or unlike a feed
// @Tags         likes
// @Produce      json
// @Param        feed_id  path      string  true  "Feed ID"
// @Success      200      {object}  response.DocBaseResponse
// @Failure      401         {object}  response.ErrorResponse
// @Failure      404         {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /feeds/{feed_id}/like [post]
func (h *LikeHandler) ToggleFeedLike(c *echo.Context) error {
	ctx := c.Request().Context()

	user := utils.RequireUser(c)
	userID := user.ID

	feedId := c.Param("feed_id")
	if feedId == "" {
		return response.BadRequest(c, h.logger, "no feed id", nil)
	}

	if err := utils.ValidateUUID(feedId); err != nil {
		return response.BadRequest(c, h.logger, "invalid feed id", err)
	}

	if err := h.feedService.CheckIfFeedExist(ctx, feedId); err != nil {
		if errors.Is(err, feed.ErrFeedNotFound) {
			return response.NotFound(c, h.logger, "feed not found", nil)
		}
		return response.InternalError(c, h.logger, err)
	}

	if err := h.service.ToggleFeedLike(ctx, userID, feedId); err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.OKMessage(c, "feed like toggled")
}

// ToggleCommentLike godoc
// @Summary      Toggle comment like
// @Description  Like or unlike a comment
// @Tags         likes
// @Produce      json
// @Param        comment_id  path      string  true  "Comment ID"
// @Success      200      {object}  response.DocBaseResponse
// @Failure      401         {object}  response.ErrorResponse
// @Failure      404         {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /comments/{comment_id}/like [post]
func (h *LikeHandler) ToggleCommentLike(c *echo.Context) error {
	ctx := c.Request().Context()

	user := utils.RequireUser(c)
	userID := user.ID

	commentId := c.Param("comment_id")
	if commentId == "" {
		return response.BadRequest(c, h.logger, "no comment id", nil)
	}

	if err := utils.ValidateUUID(commentId); err != nil {
		return response.BadRequest(c, h.logger, "invalid comment id", err)
	}

	if err := h.commentService.CheckIfCommentExist(ctx, commentId); err != nil {
		if errors.Is(err, comment.ErrCommentNotFound) {
			return response.NotFound(c, h.logger, "comment not found", nil)
		}
		return response.InternalError(c, h.logger, err)
	}

	if err := h.service.ToggleCommentLike(ctx, userID, commentId); err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.OKMessage(c, "comment like toggled")
}

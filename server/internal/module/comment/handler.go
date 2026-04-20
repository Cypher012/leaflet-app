package comment

import (
	"errors"
	"log/slog"
	"server/internal/module/feed"
	"server/internal/platform/config"
	"server/internal/shared/response"
	"server/internal/shared/utils"

	"github.com/labstack/echo/v5"
)

type CommentHandler struct {
	service     *CommentService
	feedService *feed.FeedService
	logger      *slog.Logger
	cfg         config.AppConfig
}

func NewCommentHandler(service *CommentService, feedService *feed.FeedService, logger *slog.Logger, cfg config.AppConfig) *CommentHandler {
	return &CommentHandler{service, feedService, logger, cfg}
}

// GetFeedsComment godoc
// @Summary Get feeds comment
// @Tags comments
// @Accept json
// @Produce json
// @Param	feed_id	path	string	true	"Feed Id"
// @Param   cursor query string false "Paginated cursor"
// @Success      200     {object}  DocCommentsResponse
// @Failure      400     {object}  response.ErrorResponse
// @Failure      401     {object}  response.ErrorResponse
// @Failure      500     {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /feeds/{feed_id}/comments [get]
func (h *CommentHandler) GetFeedsComments(c *echo.Context) error {
	ctx := c.Request().Context()

	var viewerID string

	if user, ok := utils.ContextUser(c); ok {
		viewerID = user.ID
	}

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

	cursor := c.QueryParam("cursor")
	comments, meta, err := h.service.FetchComments(ctx, viewerID, feedId, cursor)
	if err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.Paginated(c, "Comments fetched", comments, meta)
}

// PostComment godoc
// @Summary Create comment
// @Tags comments
// @Accept json
// @Produce json
// @Param	feed_id	path	string	true	"Feed Id"
// @Param	comment	body	NewComment	true	"Post Comment"
// @Success      200     {object}  response.DocBaseResponse
// @Failure      400     {object}  response.ErrorResponse
// @Failure      401     {object}  response.ErrorResponse
// @Failure      500     {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /feeds/{feed_id}/comments [post]
func (h *CommentHandler) PostComment(c *echo.Context) error {
	ctx := c.Request().Context()

	user := utils.RequireUser(c)
	userID := user.ID

	feedId := c.Param("feed_id")
	if feedId == "" {
		return response.BadRequest(c, h.logger, "no feed id", nil)
	}

	if err := h.feedService.CheckIfFeedExist(ctx, feedId); err != nil {
		if errors.Is(err, feed.ErrFeedNotFound) {
			return response.NotFound(c, h.logger, "feed not found", nil)
		}
		return response.InternalError(c, h.logger, err)
	}

	req := new(NewComment)
	if err := c.Bind(req); err != nil {
		return response.BadRequest(c, h.logger, "invalid body", err)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	err := h.service.PostFeedComment(ctx, userID, feedId, req.Content)
	if err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.OKMessage(c, "comment posted successfully")
}

// PostCommentReply godoc
// @Summary Reply comment
// @Tags comments
// @Accept json
// @Produce json
// @Param	comment_id	path	string	true	"Comment Id"
// @Param	comment	body	NewComment	true	"Post Comment"
// @Success      200     {object}  response.DocBaseResponse
// @Failure      400     {object}  response.ErrorResponse
// @Failure      401     {object}  response.ErrorResponse
// @Failure      500     {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /comments/{comment_id}/replies [post]
func (h *CommentHandler) ReplyComment(c *echo.Context) error {
	ctx := c.Request().Context()

	user := utils.RequireUser(c)
	userID := user.ID

	commentId := c.Param("comment_id")
	if commentId == "" {
		return response.BadRequest(c, h.logger, "no comment id", nil)
	}

	if err := h.service.CheckIfCommentExist(ctx, commentId); err != nil {
		if errors.Is(err, ErrCommentNotFound) {
			return response.NotFound(c, h.logger, "comment not found", nil)
		}
		return response.InternalError(c, h.logger, err)
	}

	req := new(NewComment)
	if err := c.Bind(req); err != nil {
		return response.BadRequest(c, h.logger, "invalid body", err)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	err := h.service.ReplyComment(ctx, userID, commentId, req.Content)
	if err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.OKMessage(c, "comment replied successfully")
}

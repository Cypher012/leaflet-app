package feed

import (
	"errors"
	"log/slog"
	"server/internal/platform/config"
	"server/internal/shared/response"
	"server/internal/shared/utils"

	"github.com/labstack/echo/v5"
)

type FeedHandler struct {
	service *FeedService
	logger  *slog.Logger
	cfg     config.AppConfig
}

func NewFeedHandler(service *FeedService, logger *slog.Logger, cfg config.AppConfig) *FeedHandler {
	return &FeedHandler{service, logger, cfg}
}

// PublicFeeds godoc
// @Summary      List public feeds
// @Description  Returns a paginated list of public feeds
// @Tags         feeds
// @Produce      json
// @Param        cursor  query     string  false  "Pagination cursor"
// @Success      200     {object}  DocFeedsResponse
// @Failure      500     {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /feeds [get]
func (h *FeedHandler) PublicFeeds(c *echo.Context) error {
	ctx := c.Request().Context()
	cursor := c.QueryParam("cursor")

	var viewerID string

	if user, ok := utils.ContextUser(c); ok {
		viewerID = user.ID
	}

	h.logger.Info("user", slog.String("user_id", viewerID))

	feeds, meta, err := h.service.FetchFeeds(ctx, cursor, viewerID)
	if err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.Paginated(c, "Feeds fetched", feeds, meta)
}

// PublicFeed godoc
// @Summary      Get a single public feed
// @Description  Returns a single feed by ID
// @Tags         feeds
// @Produce      json
// @Param        id   path      string  true  "Feed ID"
// @Success      200  {object}  DocFeedResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /feeds/{id} [get]
func (h *FeedHandler) PublicFeed(c *echo.Context) error {
	feedId := c.Param("id")

	if feedId == "" {
		return response.BadRequest(c, h.logger, "No feed id provided", nil)
	}

	var viewerID string

	if user, ok := utils.ContextUser(c); ok {
		viewerID = user.ID
	}

	feed, err := h.service.FetchFeed(c.Request().Context(), feedId, viewerID)
	if err != nil {
		if errors.Is(err, ErrFeedNotFound) {
			return response.NotFound(c, h.logger, "Feed not found", err)
		}
		return response.InternalError(c, h.logger, err)
	}

	return response.OK(c, "Feed fetched", feed)
}

// CreateFeed godoc
// @Summary      Create a feed
// @Description  Creates a new feed for the authenticated user
// @Tags         feeds
// @Accept       json
// @Produce      json
// @Param        body  body  CreateFeedReq  true  "Feed payload"
// @Success 201 {object} response.DocBaseResponse
// @Failure      400   {object}  response.ErrorResponse
// @Failure      500   {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /feeds [post]
func (h *FeedHandler) CreateFeed(c *echo.Context) error {
	user := utils.RequireUser(c)
	userID := user.ID

	h.logger.Info("create feed", "userID", userID)

	req := new(CreateFeedReq)
	if err := c.Bind(req); err != nil {
		return response.BadRequest(c, h.logger, "Invalid request body", err)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if err := h.service.CreateFeed(c.Request().Context(), userID, *req); err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.Created(c, "Feed created successfully")
}

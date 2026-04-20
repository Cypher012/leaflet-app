package profile

import (
	"context"
	"errors"
	"log/slog"
	"server/internal/module/auth"
	"server/internal/platform/config"
	"server/internal/shared/response"
	"server/internal/shared/utils"

	"github.com/labstack/echo/v5"
)

type ProfileHandler struct {
	service     *ProfileService
	authService *auth.AuthService
	logger      *slog.Logger
	cfg         config.AppConfig
}

func NewProfileHandler(service *ProfileService, authService *auth.AuthService, logger *slog.Logger, cfg config.AppConfig) *ProfileHandler {
	return &ProfileHandler{service, authService, logger, cfg}
}

// UserProfile godoc
// @Summary      Get user profile
// @Description  Returns the public profile of a user by username
// @Tags         profile
// @Produce      json
// @Param        username  path      string  true  "Username"
// @Success      200       {object}  DocUserProfileResponse
// @Failure      400       {object}  response.ErrorResponse
// @Failure      404       {object}  response.ErrorResponse
// @Failure      500       {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /profile/{username} [get]
func (h *ProfileHandler) UserProfile(c *echo.Context) error {
	ctx := c.Request().Context()

	username := c.Param("username")
	if username == "" {
		return response.BadRequest(c, h.logger, "username not provided", nil)
	}

	user, err := h.service.GetUserFromUsername(ctx, username)
	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return response.NotFound(c, h.logger, "user not found", nil)
		}
		return response.InternalError(c, h.logger, err)
	}

	return response.OK(c, "user profile fetched", user)
}

// UserProfileOverviews godoc
// @Summary      Get user activity overview
// @Description  Returns a paginated list of the user's activity
// @Tags         profile
// @Produce      json
// @Param        username  path     string  true  "Username"
// @Param        cursor  query     string  false  "Pagination cursor"
// @Success      200     {object}  DocUserActivityResponse
// @Failure      400     {object}  response.ErrorResponse
// @Failure      404     {object}  response.ErrorResponse
// @Failure      500     {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /profile/{username}/overview [get]
func (h *ProfileHandler) UserProfileOverview(c *echo.Context) error {
	ctx := c.Request().Context()

	userID, err := h.resolveUserID(ctx, c)
	if err != nil {
		return err
	}

	var viewerID string

	if user, ok := utils.ContextUser(c); ok {
		viewerID = user.ID
	}

	cursor := c.QueryParam("cursor")

	activity, meta, err := h.service.GetUserActivity(ctx, userID, viewerID, cursor)
	if err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.Paginated(c, "user activity fetched", activity, meta)
}

// UserProfilePosts godoc
// @Summary      Get user feeds
// @Description  Returns a paginated list of the user's feeds
// @Tags         profile
// @Produce      json
// @Param        username  path     string  true  "Username"
// @Param        cursor  query     string  false  "Pagination cursor"
// @Success      200     {object}  DocUserPublicFeedResponse
// @Failure      400     {object}  response.ErrorResponse
// @Failure      404     {object}  response.ErrorResponse
// @Failure      500     {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /profile/{username}/feeds [get]
func (h *ProfileHandler) UserProfileFeeds(c *echo.Context) error {
	ctx := c.Request().Context()

	userID, err := h.resolveUserID(ctx, c)
	if err != nil {
		return err
	}

	var viewerID string

	if user, ok := utils.ContextUser(c); ok {
		viewerID = user.ID
	}

	cursor := c.QueryParam("cursor")

	activity, meta, err := h.service.GetUserFeeds(ctx, userID, viewerID, cursor)
	if err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.Paginated(c, "user feeds fetched", activity, meta)
}

// UserProfileComments godoc
// @Summary      Get user comments
// @Description  Returns a paginated list of the user's comments
// @Tags         profile
// @Produce      json
// @Param        username  path     string  true  "Username"
// @Param        cursor  query     string  false  "Pagination cursor"
// @Success      200     {object}  DocUserFeedCommentResponse
// @Failure      400     {object}  response.ErrorResponse
// @Failure      404     {object}  response.ErrorResponse
// @Failure      500     {object}  response.ErrorResponse
// @Security     SessionAuth
// @Router       /profile/{username}/comments [get]
func (h *ProfileHandler) UserProfileComments(c *echo.Context) error {
	ctx := c.Request().Context()

	userID, err := h.resolveUserID(ctx, c)
	if err != nil {
		return err
	}

	var viewerID string

	if user, ok := utils.ContextUser(c); ok {
		viewerID = user.ID
	}

	cursor := c.QueryParam("cursor")

	activity, meta, err := h.service.GetUserComments(ctx, userID, viewerID, cursor)
	if err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.Paginated(c, "user comments fetched", activity, meta)
}

func (h *ProfileHandler) resolveUserID(ctx context.Context, c *echo.Context) (string, error) {
	username := c.Param("username")
	if username == "" {
		return "", response.BadRequest(c, h.logger, "username not provided", nil)
	}

	userID, err := h.authService.GetUserIDFromUsername(ctx, username)
	if err != nil {
		if errors.Is(err, auth.ErrUsernameNotFound) {
			return "", response.NotFound(c, h.logger, "username not found", nil)
		}
		return "", response.InternalError(c, h.logger, err)
	}

	return userID, nil
}

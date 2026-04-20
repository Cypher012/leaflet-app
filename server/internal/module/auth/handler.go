package auth

import (
	"log/slog"
	"net/http"
	"server/internal/platform/config"
	"server/internal/shared/response"
	"server/internal/shared/utils"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	service *AuthService
	logger  *slog.Logger
	cfg     config.AppConfig
}

func NewAuthHandler(service *AuthService, logger *slog.Logger, cfg config.AppConfig) *AuthHandler {
	return &AuthHandler{service, logger, cfg}
}

// OAuthLogin godoc
// @Summary OAuth login
// @Description Redirects user to OAuth provider (Google, GitHub, etc.)
// @Tags auth
// @Produce json
// @Param provider path string true "OAuth provider (e.g. google, github)"
// @Success 302 {string} string "Redirect to frontend after successful login"
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/{provider} [get]
func (h *AuthHandler) OAuthLogin(c *echo.Context) error {
	provider := c.Param("provider")
	if provider == "" {
		return response.BadRequest(c, h.logger, "provider not specified", nil)
	}

	q := c.QueryParams()
	q.Add("provider", provider)
	c.Request().URL.RawQuery = q.Encode()

	req, res := c.Request(), c.Response()

	if _, err := gothic.CompleteUserAuth(res, req); err == nil {
		return response.Redirect(c, h.cfg.FrontendURL)
	}

	gothic.BeginAuthHandler(res, req)

	return nil
}

// OAuthCallback godoc
// @Summary OAuth callback
// @Description Handles OAuth provider callback, creates session, and sets cookie
// @Tags auth
// @Produce plain
// @Param provider path string true "OAuth provider (e.g. google, github)"
// @Success 302 {string} string "Redirect to frontend after successful authentication"
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/{provider}/callback [get]
func (h *AuthHandler) OAuthCallback(c *echo.Context) error {
	req, res := c.Request(), c.Response()
	ctx := req.Context()

	gothUser, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		return response.BadRequest(c, h.logger, "oauth authentication failed", err)
	}

	if gothUser.Provider == "github" && gothUser.Email == "" {
		email, err := utils.FetchGitHubEmail(gothUser.AccessToken)
		if err != nil {
			return response.BadRequest(c, h.logger, "error fetching github email", err)
		}
		gothUser.Email = email
	}

	dbUserId, err := h.service.OauthCallback(ctx, gothUser)
	if err != nil {
		return response.InternalError(c, h.logger, err)
	}

	ip := c.RealIP()
	userAgent := req.UserAgent()

	session, err := h.service.CreateSession(ctx, dbUserId, ip, userAgent)
	if err != nil {
		return response.InternalError(c, h.logger, err)
	}

	c.SetCookie(&http.Cookie{
		Name:     "leaflet_sid",
		Value:    session.Token,
		Path:     "/",
		Domain:   "leaflet-dev.com",
		Expires:  session.ExpiresAt,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	})

	return c.Redirect(http.StatusFound, h.cfg.FrontendURL)

}

// GetSession godoc
// @Summary Get user session
// @Description Returns the currently authenticated user session
// @Tags auth
// @Produce json
// @Security SessionAuth
// @Success 200 {object} DocSessionResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /auth/session [get]
func (h *AuthHandler) GetCurrentSession(c *echo.Context) error {
	ctx := c.Request().Context()

	cookie, err := c.Cookie("leaflet_sid")
	if err != nil {
		return response.UnauthorizedError(c, h.logger, nil)
	}

	session, err := h.service.GetSession(ctx, cookie.Value)
	if err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.OK(c, "session fetched successfully", session)
}

// GetMe godoc
// @Summary Get current user
// @Description Returns basic profile information of the currently authenticated user
// @Tags auth
// @Produce json
// @Security SessionAuth
// @Success 200 {object} DocMinimalProfileResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/me [get]
func (h *AuthHandler) GetMe(c *echo.Context) error {
	ctx := c.Request().Context()

	cookie, err := c.Cookie("leaflet_sid")
	if err != nil {
		return response.UnauthorizedError(c, h.logger, nil)
	}

	user, err := h.service.GetUserFromToken(ctx, cookie.Value)
	if err != nil {
		return response.InternalError(c, h.logger, err)
	}

	return response.OK(c, "current user fetched successfully", user)
}

// Logout godoc
// @Summary Logout user
// @Description Deletes user session and clears session cookie
// @Tags auth
// @Produce json
// @Security SessionAuth
// @Success 204 {string} string "No Content"
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *echo.Context) error {
	cookie, err := c.Cookie("leaflet_sid")
	if err != nil {
		return response.UnauthorizedError(c, h.logger, nil)
	}

	if err := h.service.Logout(c.Request().Context(), cookie.Value); err != nil {
		return response.InternalError(c, h.logger, err)
	}

	c.SetCookie(&http.Cookie{
		Name:     "leaflet_sid",
		Value:    "",
		Domain:   "leaflet-dev.com",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		Path:     "/",
	})

	return response.OKMessage(c, "Signed out successfully")
}

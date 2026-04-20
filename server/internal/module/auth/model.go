package auth

import (
	"errors"
	db "server/internal/platform/db"
	"server/internal/shared/response"
	"server/internal/shared/types"

	"time"
)

type DocUserResponse response.Response[UserResponse]
type DocMinimalProfileResponse response.Response[MinimalUserProfile]
type DocSessionResponse response.Response[SessionResponse]

// error type
var (
	ErrUsernameNotFound = errors.New("user not found")
	ErrUserNotFound     = errors.New("user not found")
	ErrSessionNotFound  = errors.New("session not found")
)

type CreateSessionParams struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
	IpAddress string
	UserAgent string
}

type UserResponse struct {
	ID            string    `json:"id"`
	Email         string    `json:"email,omitempty"`
	Fullname      string    `json:"fullname"`
	Username      string    `json:"username"`
	Bio           string    `json:"bio,omitempty"`
	EmailVerified bool      `json:"email_verified,omitempty"`
	AvatarUrl     string    `json:"avatar_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func NormalizeUserResponse(u types.User) UserResponse {
	return UserResponse{
		ID:            u.ID,
		Email:         u.Email,
		Fullname:      u.Fullname,
		Username:      u.Username,
		Bio:           u.Bio,
		EmailVerified: u.EmailVerified,
		AvatarUrl:     u.AvatarUrl,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}
}

type SessionResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"-"`
	IpAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NormalizeSessionResponse(s types.Session) SessionResponse {
	return SessionResponse{
		ID:        s.ID,
		UserID:    s.UserID,
		Token:     s.Token,
		IpAddress: s.IpAddress,
		UserAgent: s.UserAgent,
		ExpiresAt: s.ExpiresAt,
	}
}

type MinimalUserProfile struct {
	Fullname  string `json:"fullname"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}

func NormalizeMinimalProfile(u db.GetUserByTokenRow) MinimalUserProfile {
	return MinimalUserProfile{
		Fullname:  u.Fullname,
		Username:  u.Username,
		AvatarUrl: u.AvatarUrl.String,
	}
}

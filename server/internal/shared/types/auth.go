package types

import (
	"server/internal/platform/db"
	"time"
)

type Session struct {
	ID        string
	UserID    string
	Token     string
	IpAddress string
	UserAgent string
	ExpiresAt time.Time
}

type User struct {
	ID            string
	Email         string
	Fullname      string
	Username      string
	Bio           string
	EmailVerified bool
	AvatarUrl     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NormalizeUser(u db.User) User {
	return User{
		ID:            u.ID.String(),
		Email:         u.Email,
		Fullname:      u.Fullname,
		Username:      u.Username,
		Bio:           u.Bio.String,
		EmailVerified: u.EmailVerified,
		AvatarUrl:     u.AvatarUrl.String,
		CreatedAt:     u.CreatedAt.Time,
		UpdatedAt:     u.UpdatedAt.Time,
	}
}

func NormalizeSession(s db.Session) Session {
	return Session{
		ID:        s.ID.String(),
		UserID:    s.UserID.String(),
		Token:     s.Token,
		IpAddress: s.IpAddress.String,
		UserAgent: s.UserAgent.String,
		ExpiresAt: s.ExpiresAt.Time,
	}
}

func NormalizeSessions(sessions []db.Session) []Session {
	result := make([]Session, 0, len(sessions))

	for _, s := range sessions {
		result = append(result, NormalizeSession(s))
	}

	return result
}

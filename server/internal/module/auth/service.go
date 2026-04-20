package auth

import (
	"context"
	"errors"
	"fmt"
	"server/internal/shared/utils"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/markbates/goth"
)

type AuthService struct {
	repo *AuthRepository
}

func NewAuthService(repo *AuthRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) OauthCallback(ctx context.Context, gothUser goth.User) (dbUserID string, err error) {
	userID, err := s.repo.OAuthFindOrCreateUserWithAccount(ctx, gothUser)
	if err != nil {
		return "", fmt.Errorf("OAuthCallback: %w", err)
	}

	return userID, nil
}

func (s *AuthService) CreateSession(ctx context.Context, userID string, ip, userAgent string) (session SessionResponse, err error) {
	token, err := utils.GenerateSessionToken()
	if err != nil {
		return SessionResponse{}, err
	}

	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	dbSession, err := s.repo.CreateSession(ctx, CreateSessionParams{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
		IpAddress: ip,
		UserAgent: userAgent,
	})

	if err != nil {
		return SessionResponse{}, err
	}

	return NormalizeSessionResponse(dbSession), nil
}
func (s *AuthService) GetSession(ctx context.Context, token string) (SessionResponse, error) {
	dbSession, err := s.repo.GetSessionByToken(ctx, token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return SessionResponse{}, ErrSessionNotFound
		}
		return SessionResponse{}, err
	}
	return NormalizeSessionResponse(dbSession), nil
}

func (s *AuthService) GetUserFromToken(ctx context.Context, token string) (MinimalUserProfile, error) {
	user, err := s.repo.GetUserByToken(ctx, token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return MinimalUserProfile{}, ErrUserNotFound
		}
		return MinimalUserProfile{}, err
	}
	return user, nil
}

func (s *AuthService) GetUserIDFromUsername(ctx context.Context, username string) (string, error) {
	id, err := s.repo.GetUserIDByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrUsernameNotFound
		}
		return "", err
	}
	return id, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	return s.repo.DeleteSessionByToken(ctx, token)
}

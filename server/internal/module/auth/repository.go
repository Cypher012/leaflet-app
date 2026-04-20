package auth

import (
	"context"
	"errors"

	db "server/internal/platform/db"
	"server/internal/shared/types"
	"server/internal/shared/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/markbates/goth"
)

type AuthRepository struct {
	q    *db.Queries
	conn *pgxpool.Pool
}

func NewAuthRepository(conn *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		q:    db.New(conn),
		conn: conn,
	}
}

// ── User ─────────────────────────────────────────────────────────────────────

func (r *AuthRepository) CreateUser(ctx context.Context, params db.CreateUserParams) (types.User, error) {
	dbUser, err := r.q.CreateUser(ctx, params)
	if err != nil {
		return types.User{}, err
	}
	return types.NormalizeUser(dbUser), nil
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (types.User, error) {
	dbUser, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return types.User{}, err
	}
	return types.NormalizeUser(dbUser), nil
}

func (r *AuthRepository) GetUserByID(ctx context.Context, id string) (types.User, error) {
	uuid, err := utils.ConvertIdToPgUUID(&id)
	if err != nil {
		return types.User{}, err
	}

	dbUser, err := r.q.GetUserByID(ctx, uuid)
	if err != nil {
		return types.User{}, err
	}
	return types.NormalizeUser(dbUser), nil
}

func (r *AuthRepository) GetUserIDByUsername(ctx context.Context, username string) (string, error) {
	uuid, err := r.q.GetUserIDByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	id, err := utils.PgUUIDToString(uuid)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *AuthRepository) GetUserByToken(ctx context.Context, username string) (MinimalUserProfile, error) {
	dbUser, err := r.q.GetUserByToken(ctx, username)
	if err != nil {
		return MinimalUserProfile{}, err
	}
	return NormalizeMinimalProfile(dbUser), nil
}

// ── Session ───────────────────────────────────────────────────────────────────

func (r *AuthRepository) CreateSession(ctx context.Context, params CreateSessionParams) (types.Session, error) {
	uuid, err := utils.ConvertIdToPgUUID(&params.UserID)
	if err != nil {
		return types.Session{}, err
	}
	dbSession, err := r.q.CreateSession(ctx, db.CreateSessionParams{
		UserID:    uuid,
		Token:     params.Token,
		ExpiresAt: utils.PgTimestamptz(params.ExpiresAt),
		IpAddress: utils.PgText(params.IpAddress),
		UserAgent: utils.PgText(params.UserAgent),
	})
	if err != nil {
		return types.Session{}, err
	}
	return types.NormalizeSession(dbSession), nil
}

func (r *AuthRepository) GetSessionByToken(ctx context.Context, token string) (types.Session, error) {
	session, err := r.q.GetSessionByToken(ctx, token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return types.Session{}, ErrSessionNotFound
		}
		return types.Session{}, err
	}

	return types.NormalizeSession(session), nil

}

func (r *AuthRepository) ListUserSessions(ctx context.Context, userID string) ([]types.Session, error) {
	uuid, err := utils.ConvertIdToPgUUID(&userID)
	if err != nil {
		return nil, err
	}
	sessions, err := r.q.ListUserSessions(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return types.NormalizeSessions(sessions), nil
}

func (r *AuthRepository) RevokeSession(ctx context.Context, token string) error {
	return r.q.RevokeSession(ctx, token)
}

func (r *AuthRepository) RevokeAllUserSessions(ctx context.Context, userID string) error {
	uuid, err := utils.ConvertIdToPgUUID(&userID)
	if err != nil {
		return err
	}
	return r.q.RevokeAllUserSessions(ctx, uuid)
}

func (r *AuthRepository) DeleteSessionByID(ctx context.Context, id string) error {
	uuid, err := utils.ConvertIdToPgUUID(&id)
	if err != nil {
		return err
	}
	return r.q.DeleteSessionByID(ctx, uuid)
}

func (r *AuthRepository) DeleteSessionByToken(ctx context.Context, token string) error {
	return r.q.DeleteSessionByToken(ctx, token)
}

func (r *AuthRepository) DeleteExpiredSessions(ctx context.Context) error {
	return r.q.DeleteExpiredSessions(ctx)
}

// ── Account ───────────────────────────────────────────────────────────────────

func (r *AuthRepository) CreateAccount(ctx context.Context, params db.CreateAccountParams) (db.Account, error) {
	return r.q.CreateAccount(ctx, params)
}

func (r *AuthRepository) GetAccountByProvider(ctx context.Context, params db.GetAccountByProviderParams) (db.Account, error) {
	return r.q.GetAccountByProvider(ctx, params)
}

// ── Transactions ───────────────────────────────────────────────────────────────────

func (r *AuthRepository) OAuthFindOrCreateUserWithAccount(ctx context.Context, gothUser goth.User) (string, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	qtx := r.q.WithTx(tx)

	existingUser, userErr := qtx.GetUserByEmail(ctx, gothUser.Email)

	var userID pgtype.UUID

	if errors.Is(userErr, pgx.ErrNoRows) {
		dbUser, err := qtx.CreateUser(ctx, db.CreateUserParams{
			Email:     gothUser.Email,
			Fullname:  gothUser.Name,
			AvatarUrl: utils.PgText(gothUser.AvatarURL),
			Username:  utils.GenerateUsername(gothUser.Name),
		})
		if err != nil {
			return "", err
		}
		userID = dbUser.ID
	} else if userErr != nil {
		return "", userErr
	} else {
		userID = existingUser.ID
	}

	_, err = qtx.CreateAccount(ctx, db.CreateAccountParams{
		UserID:               userID,
		ProviderID:           gothUser.Provider,
		AccountID:            gothUser.UserID,
		AccessToken:          utils.PgText(gothUser.AccessToken),
		RefreshToken:         utils.PgText(gothUser.RefreshToken),
		PasswordHash:         utils.PgText(""),
		AccessTokenExpiresAt: utils.PgTimestamptz(gothUser.ExpiresAt),
	})
	if err != nil {
		return "", err
	}

	userIDStr, err := utils.PgUUIDToString(userID)
	if err != nil {
		return "", err
	}

	return userIDStr, tx.Commit(ctx)
}

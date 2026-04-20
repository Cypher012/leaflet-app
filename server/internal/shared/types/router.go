package types

import (
	"log/slog"
	"server/internal/platform/config"
	"server/internal/platform/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RouterDeps struct {
	Logger    *slog.Logger
	Conn      *pgxpool.Pool
	Config    config.AppConfig
	S3Storage *storage.Storage
}

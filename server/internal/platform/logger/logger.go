package logger

import (
	"log"
	"log/slog"
	"os"
	"server/internal/platform/config"
)

func NewLogger() *slog.Logger {
	env, err := config.GetEnv("APP_ENV")
	if err != nil {
		log.Fatal(err)
	}

	level := slog.LevelDebug
	if env == "production" {
		level = slog.LevelInfo
	}

	handlerOpts := &slog.HandlerOptions{
		Level: level,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, handlerOpts)).With(
		slog.String("service", "leaflet-server"),
		slog.String("env", env),
	)

	return logger
}

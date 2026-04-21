package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"server/internal/platform/config"
	"server/internal/platform/db"
	"server/internal/platform/logger"
	"server/internal/platform/postgres"
	"server/internal/platform/storage"
	mw "server/internal/shared/middleware"
	"server/internal/shared/types"
	vd "server/internal/shared/validator"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

// @title Leaflet API
// @version 1.0
// @description Leaflet discussion forum platform API
// @BasePath /api

// @securityDefinitions.apikey SessionAuth
// @in cookie
// @name leaflet_sid
func main() {
	logger := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("config error", slog.Any("error", err))
		os.Exit(1)
	}

	conn := postgres.NewDB(cfg.DatabaseURL)
	defer conn.Close()

	s3Storage := storage.New(cfg)

	// start cleanup job
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for range ticker.C {
			db.New(conn).DeleteExpiredSessions(context.Background())
		}
	}()

	gothic.Store = sessions.NewCookieStore([]byte(cfg.SessionSecret))

	goth.UseProviders(
		github.New(
			cfg.GithubClientID,
			cfg.GithubClientSecret,
			cfg.GithubCallbackURL,
			"user:email",
		),
		google.New(
			cfg.GoogleClientID,
			cfg.GoogleClientSecret,
			cfg.GoogleCallbackURL,
			"email", "profile",
		),
	)

	e := echo.New()
	e.Use(mw.LoggerMiddleware(logger))

	e.Validator = &vd.CustomValidator{
		Validator: validator.New(),
	}

	e.Use(middleware.ContextTimeout(10 * time.Second))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{cfg.FrontendURL},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
	}))

	NewRouter(e, types.RouterDeps{
		Logger:    logger,
		Conn:      conn,
		Config:    cfg,
		S3Storage: s3Storage,
	})

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	sc := echo.StartConfig{Address: ":" + cfg.Port}

	logger.Info("server started", slog.String("url", cfg.BackendURL))

	if err := sc.Start(ctx, e); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(err.Error())
	}
}

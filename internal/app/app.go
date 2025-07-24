package app

import (
	"context"
	"electronic-document-service/internal/config"
	"electronic-document-service/internal/db"
	"electronic-document-service/internal/server"
	"fmt"
	"log/slog"
)

// App contains all application dependency and launch http server.
type App struct {
	Server *server.Server
}

func New(log *slog.Logger, cfg config.Config) (*App, error) {
	_, err := db.NewConn(cfg.Storage)
	if err != nil {
		return nil, fmt.Errorf("can't connect to storage: %w", err)
	}

	srv := server.New(log, cfg.Server)
	srv.ConfigureRoutes()

	return &App{
		Server: srv,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	if err := a.Server.Run(ctx); err != nil {
		return fmt.Errorf("app run: %w", err)
	}

	return nil
}

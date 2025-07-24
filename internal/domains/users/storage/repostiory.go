package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Users struct {
	log *slog.Logger
	DB  *sqlx.DB
}

func NewUsersRepo(log *slog.Logger, db *sqlx.DB) *Users {
	return &Users{log: log, DB: db}
}

func (r *Users) Create(ctx context.Context) {

}

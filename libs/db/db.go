package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	_ "github.com/jackc/pgx/v4/stdlib" // for pgx driver
	"time"
)

type Config interface {
	GetDSN() string
	GetMaxOpenConns() int
	GetMaxIdleConns() int
	GetConnMaxIdleTime() time.Duration
	GetConnMaxLifeTime() time.Duration
}

func ConnectDB(ctx context.Context, cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", cfg.GetDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.GetMaxOpenConns())
	db.SetMaxIdleConns(cfg.GetMaxIdleConns())
	db.SetConnMaxIdleTime(cfg.GetConnMaxIdleTime())
	db.SetConnMaxLifetime(cfg.GetConnMaxLifeTime())

	err = db.PingContext(ctx)
	if err != nil {
		err = errors.Wrap(err, "db.PingContext()")
		return nil, err
	}

	return db, nil
}

func CloseDB(db *sqlx.DB) {
	_ = db.Close()
}

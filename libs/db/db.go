package db

import (
	"github.com/jmoiron/sqlx"

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

func ConnectDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", cfg.GetDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.GetMaxOpenConns())
	db.SetMaxIdleConns(cfg.GetMaxIdleConns())
	db.SetConnMaxIdleTime(cfg.GetConnMaxIdleTime())
	db.SetConnMaxLifetime(cfg.GetConnMaxLifeTime())

	return db, nil
}

func CloseDB(db *sqlx.DB) {
	_ = db.Close()
}

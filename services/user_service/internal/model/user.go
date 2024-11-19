package model

import (
	"database/sql"
	"time"
)

type User struct {
	UID          string         `db:"uid"`
	Name         sql.NullString `db:"name"`
	Email        sql.NullString `db:"email"`
	PictureURL   sql.NullString `db:"picture_url"`
	Locale       sql.NullString `db:"locale"`
	PremiumType  int64          `db:"premium_type"`
	IsAnonymous  bool           `db:"is_anonymous"`
	IsEnabled    bool           `db:"is_enabled"`
	LaunchCount  int64          `db:"launch_count"`
	AppVersion   string         `db:"app_version"`
	BuildVersion sql.NullString `db:"build_version"`
	DeviceInfo   sql.NullString `db:"device_info"`
	FBCreatedAt  sql.NullTime   `db:"fb_created_at"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}

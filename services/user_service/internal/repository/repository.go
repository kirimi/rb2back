package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kirimi/rb2backend/libs/db"
	"github.com/kirimi/rb2backend/user_service/internal/model"
)

type Repository interface {
	GetUserByUID(ctx context.Context, uid string) (model.User, error)
}

type repo struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repo{db: db}
}

func (r *repo) GetUserByUID(ctx context.Context, uid string) (model.User, error) {
	user := model.User{}
	err := db.WithTx(
		ctx,
		r.db,
		func(ctx context.Context, tx *sqlx.Tx) error {
			return tx.GetContext(ctx, &user, "SELECT * FROM public.users WHERE uid=$1", uid)
		},
	)

	return user, err
}

package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type withTxFunc func(ctx context.Context, tx *sqlx.Tx) error

func WithTx(ctx context.Context, db *sqlx.DB, fn withTxFunc) error {
	t, err := db.BeginTxx(ctx, nil)

	if err != nil {
		return errors.Wrap(err, "db.BeginTxx()")
	}

	if err = fn(ctx, t); err != nil {
		if errRollback := t.Rollback(); errRollback != nil {
			return errors.Wrap(errRollback, "tx.Rollback()")
		}
		return errors.Wrap(err, "tx.withTxFunc()")
	}

	if errCommit := t.Commit(); errCommit != nil {
		return errors.Wrap(errCommit, "tx.Commit()")
	}

	return nil
}

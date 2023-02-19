package sql

import (
	"context"
	"fmt"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
)

func (r *Repository) WithinTransaction(ctx context.Context, fn func(context.Context, *transaction.DBTransaction) error) error {
	tx := &transaction.DBTransaction{}
	var err error
	r.log.Debug("Starting SQL transaction")
	tx.Sql, err = r.DB.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	r.log.Debug("Running SQL transaction")
	if err := fn(ctx, tx); err != nil {
		r.log.Debug("Rolling Back SQL transaction")
		if err := tx.Sql.Rollback(); err != nil {
			r.log.Warn("rollback tx: %s", err.Error())
		}
		return fmt.Errorf("run tx: %w", err)
	}
	r.log.Debug("Running SQL commit")
	if err := tx.Sql.Commit(); err != nil {
		r.log.Debug("Rolling Back SQL transaction")
		if err := tx.Sql.Rollback(); err != nil {
			r.log.Warn("rollback tx: %s", err.Error())
		}
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}

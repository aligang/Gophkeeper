package sql

import (
	"context"
	"fmt"
	"github.com/aligang/Gophkeeper/internal/server/repository/transaction"
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
		r.log.Warn("Rolling Back SQL transaction")
		if err := tx.Sql.Rollback(); err != nil {
			r.log.Crit("rollback tx: %s", err.Error())
		}
		return fmt.Errorf("run tx: %w", err)
	}
	r.log.Debug("Running SQL commit")
	if err := tx.Sql.Commit(); err != nil {
		r.log.Warn("Rolling Back SQL transaction")
		if err := tx.Sql.Rollback(); err != nil {
			r.log.Crit("rollback tx: %s", err.Error())
		}
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}

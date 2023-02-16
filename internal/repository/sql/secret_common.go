package sql

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/repository/transaction"
	"github.com/jmoiron/sqlx"
)

func (r *Repository) addSecret(ctx context.Context, secretType string, query string, args []any, tx *transaction.DBTransaction) error {
	statement, err := r.prepareStatement(ctx, query, tx)

	r.log.Debug("Executing statement to create %s within Repository: %s", secretType, args[0])
	_, err = statement.ExecContext(ctx, args...)
	if err != nil {
		r.log.Warn("Error During statement Execution for create %s secret with id:  %s", secretType, args[0])
		return err
	}
	return nil
}

func (r *Repository) updateSecret(ctx context.Context, secretType string, query string, args []any, tx *transaction.DBTransaction) error {
	statement, err := r.prepareStatement(ctx, query, tx)

	r.log.Debug("Executing statement to update %s to Repository: %+s", secretType, args[0])
	_, err = statement.ExecContext(ctx, args...)
	if err != nil {
		r.log.Warn("Error During statement Execution for updating %s secret with id:  %s", secretType, args[0])
		return err
	}
	return nil
}

func (r *Repository) deleteSecret(ctx context.Context, secretType string, query string, args []any, tx *transaction.DBTransaction) error {

	statement, err := r.prepareStatement(ctx, query, tx)
	r.log.Debug("Executing statement to delete %s secret from to Repository: %s", secretType, args[0])
	_, err = statement.ExecContext(ctx, args...)
	if err != nil {
		r.log.Warn("Error During statement Execution %s with %s", query, args[0])
		return err
	}
	return nil
}

func (r *Repository) prepareStatement(ctx context.Context, query string, tx *transaction.DBTransaction) (*sqlx.Stmt, error) {
	var err error
	var statement *sqlx.Stmt
	if tx != nil {
		statement, err = tx.Sql.PreparexContext(ctx, query)
	} else {
		statement, err = r.DB.PreparexContext(ctx, query)
	}
	if err != nil {
		r.log.Warn("Error During statement creation %s", query)
		return nil, err
	}
	return statement, nil
}

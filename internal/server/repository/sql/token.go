package sql

import (
	"context"
	"database/sql"
	"errors"
	tokenInstance "github.com/aligang/Gophkeeper/internal/common/token/instance"
	"github.com/aligang/Gophkeeper/internal/server/repository/repositoryerrors"
	"github.com/aligang/Gophkeeper/internal/server/repository/transaction"
	"github.com/jmoiron/sqlx"
)

func (r *Repository) AddToken(ctx context.Context, t *tokenInstance.Token, tx *transaction.DBTransaction) error {
	query := "INSERT INTO tokens (Id, TokenValue, Owner, IssuedAt) VALUES($1, $2, $3, $4)"
	args := []any{t.Id, t.TokenValue, t.Owner, t.IssuedAt}
	r.log.Debug("Preparing statement to create token to Repository: %+s", t.Id)
	statement, err := tx.Sql.PreparexContext(ctx, query)
	if err != nil {
		r.log.Crit("Error During statement creation %s", query)
		return err
	}
	r.log.Debug("Executing statement to create token to Repository: %+s", t.Id)
	_, err = statement.ExecContext(ctx, args...)
	if err != nil {
		r.log.Crit("Error During statement Execution for storing token with id:  %s", t.Id)
		return err
	}
	return nil
}

func (r *Repository) GetToken(ctx context.Context, tokenValue string, tx *transaction.DBTransaction) (*tokenInstance.Token, error) {
	query := "SELECT * FROM tokens WHERE TokenValue = $1"
	var args = []interface{}{tokenValue}
	r.log.Debug("Preparing statement to fetch token from Repository: ")

	a := &tokenInstance.Token{}

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
	r.log.Debug("Executing statement to get token from Repository")
	err = statement.GetContext(ctx, a, args...)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		r.log.Warn("Database response is empty")
		return nil, repositoryerrors.ErrNoContent
	case err != nil:
		r.log.Crit("Error during decoding database response: %s", err.Error())
		return nil, err
	default:
		return a, nil
	}
}

func (r *Repository) ListAccountTokens(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*tokenInstance.Token, error) {
	r.log.Debug("Preparing statement to fetch account tokens from Repository")
	query := "SELECT Id, TokenValue, Owner, IssuedAt FROM tokens where owner = $1 ORDER BY IssuedAt DESC;"
	args := []interface{}{accountID}
	return r.listTokens(ctx, query, args, tx)
}

func (r *Repository) ListTokens(ctx context.Context, tx *transaction.DBTransaction) ([]*tokenInstance.Token, error) {
	r.log.Debug("Preparing statement to fetch account tokens from Repository")
	query := "SELECT Id, TokenValue, Owner, IssuedAt  FROM tokens ORDER BY IssuedAt ASC"
	args := []interface{}{}
	return r.listTokens(ctx, query, args, tx)
}

func (r *Repository) listTokens(ctx context.Context, query string, args []any, tx *transaction.DBTransaction) ([]*tokenInstance.Token, error) {
	var tokens []*tokenInstance.Token

	var err error
	var statement *sqlx.Stmt
	if tx != nil {
		statement, err = tx.Sql.PreparexContext(ctx, query)
	} else {
		statement, err = r.DB.PreparexContext(ctx, query)
	}
	if err != nil {
		r.log.Crit("Error During statement creation %s", query)
		return nil, err
	}
	r.log.Debug("Executing statement to fetch orders from Repository")
	err = statement.SelectContext(ctx, &tokens, args...)
	if err != nil {
		r.log.Crit("Error During statement Execution %s with %s", query, args[0])
		return tokens, err
	}
	return tokens, nil
}

func (r *Repository) DeleteToken(ctx context.Context, t *tokenInstance.Token, tx *transaction.DBTransaction) error {
	r.log.Debug("Preparing statement to delete token from repository %s", t.Id)
	query := "DELETE FROM tokens WHERE Id = $1"
	var args = []interface{}{t.Id}

	var err error
	var statement *sqlx.Stmt
	if tx != nil {
		statement, err = tx.Sql.PreparexContext(ctx, query)
	} else {
		statement, err = r.DB.PreparexContext(ctx, query)
	}
	if err != nil {
		r.log.Warn("Error During statement creation %s", query)
		return err
	}
	r.log.Debug("Executing statement to delete token from to Repository: %+s", t.Id)
	_, err = statement.ExecContext(ctx, args...)
	if err != nil {
		r.log.Warn("Error During statement Execution %s with %s", query, args[0])
		return err
	}
	return nil
}

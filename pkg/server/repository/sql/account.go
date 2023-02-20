package sql

import (
	"context"
	"database/sql"
	"errors"
	accountInstance "github.com/aligang/Gophkeeper/pkg/common/account/instance"
	"github.com/aligang/Gophkeeper/pkg/server/repository/repositoryerrors"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
	"github.com/jmoiron/sqlx"
)

func (r *Repository) Register(ctx context.Context, account *accountInstance.Account, tx *transaction.DBTransaction) error {
	query := "INSERT INTO accounts (Id, CreatedAt, Login, Password, EncryptionEnabled, EncryptionKey) VALUES($1, $2, $3, $4, $5, $6)"
	args := []any{account.Id, account.CreatedAt, account.Login, account.Password, account.EncryptionEnabled, account.EncryptionKey}
	r.log.Debug("Preparing statement to register account to Repository: %+v", account.Id)
	statement, err := tx.Sql.PreparexContext(ctx, query)
	if err != nil {
		r.log.Crit("Error During statement creation %s", query)
		return err
	}
	r.log.Debug("Executing statement to register customer account to Repository: %+v", account.Id)
	_, err = statement.ExecContext(ctx, args...)
	if err != nil {
		r.log.Crit("Error During statement Execution %s with %s, *****, %s, %s, %s",
			query, args[0], args[1], args[2], args[3], args[4], args[5])
		return err
	}
	return nil
}

func (r *Repository) GetAccountByLogin(ctx context.Context, login string, tx *transaction.DBTransaction) (*accountInstance.Account, error) {
	query := "SELECT id, login, password, CreatedAt, EncryptionEnabled, EncryptionKey FROM accounts WHERE login = $1"
	var args = []interface{}{login}
	r.log.Debug("Preparing statement to fetch customer account to Repository: %s", login)
	return r.getAccount(ctx, query, args, tx)
}

func (r *Repository) GetAccountById(ctx context.Context, accountID string, tx *transaction.DBTransaction) (*accountInstance.Account, error) {
	query := "SELECT id, login, password, CreatedAt, EncryptionEnabled, EncryptionKey FROM accounts WHERE Id = $1"
	var args = []interface{}{accountID}
	r.log.Crit("Preparing statement to fetch customer account to Repository: %s", accountID)
	return r.getAccount(ctx, query, args, tx)
}

func (r *Repository) getAccount(ctx context.Context, query string, args []any, tx *transaction.DBTransaction) (*accountInstance.Account, error) {
	var err error
	var statement *sqlx.Stmt
	a := &accountInstance.Account{}

	if tx != nil {
		statement, err = tx.Sql.PreparexContext(ctx, query)
	} else {
		statement, err = r.DB.PreparexContext(ctx, query)
	}
	if err != nil {
		r.log.Crit("Error During statement creation %s", query)
		return nil, err
	}
	r.log.Debug("Executing statement to get customer account from Repository")
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

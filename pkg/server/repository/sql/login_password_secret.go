package sql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aligang/Gophkeeper/pkg/secret/instance"
	"github.com/aligang/Gophkeeper/pkg/server/repository/repositoryerrors"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
)

func (r *Repository) AddLoginPasswordSecret(ctx context.Context, s *instance.LoginPasswordSecret, tx *transaction.DBTransaction) error {
	query := "INSERT INTO login_password_secrets (Id, AccountId, CreatedAt, ModifiedAt, Login, Password) VALUES($1, $2, $3, $4, $5, $6)"
	args := []any{s.Id, s.AccountId, s.CreatedAt, s.ModifiedAt, s.Login, s.Password}

	secretType := "login-password"
	r.log.Debug("Preparing statement to create login-password secret to Repository: %+s", args[0])
	return r.addSecret(ctx, secretType, query, args, tx)
}

func (r *Repository) UpdateLoginPasswordSecret(ctx context.Context, s *instance.LoginPasswordSecret, tx *transaction.DBTransaction) error {
	query := "UPDATE login_password_secrets SET Id = $1, AccountId = $2, CreatedAt = $3, ModifiedAt = $4, Login = $5, Password = $6 WHERE Id = $7"
	args := []any{s.Id, s.AccountId, s.CreatedAt, s.ModifiedAt, s.Login, s.Password, s.Id}
	secretType := "login-password"
	r.log.Debug("Preparing statement to update login-password secret to Repository: %s", args[0])
	return r.updateSecret(ctx, secretType, query, args, tx)
}

func (r *Repository) GetLoginPasswordSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance.LoginPasswordSecret, error) {
	query := "SELECT * FROM login_password_secrets WHERE id = $1"
	var args = []interface{}{secretID}
	r.log.Debug("Preparing statement to fetch login password from Repository: ")
	statement, err := r.prepareStatement(ctx, query, tx)
	s := &instance.LoginPasswordSecret{}

	r.log.Debug("Executing statement to login password secret  from Repository")
	err = statement.GetContext(ctx, s, args...)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		r.log.Warn("Database response is empty")
		return nil, repositoryerrors.ErrNoContent
	case err != nil:
		r.log.Warn("Error during decoding database response: %s", err.Error())
		return nil, err
	default:
		return s, nil
	}
}

func (r *Repository) ListLoginPasswordSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance.LoginPasswordSecret, error) {
	r.log.Debug("Preparing statement to fetch account login-password secrets from Repository")
	query := "SELECT * FROM login_password_secrets WHERE accountId = $1"
	args := []interface{}{accountID}
	statement, err := r.prepareStatement(ctx, query, tx)

	var secrets []*instance.LoginPasswordSecret
	r.log.Debug("Executing statement to fetch login-password secret from Repository")
	err = statement.SelectContext(ctx, &secrets, args...)
	if err != nil {
		r.log.Warn("Error During statement Execution %s with %s", query, args[0])
		return secrets, err
	}
	return secrets, nil
}

func (r *Repository) DeleteLoginPasswordSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error {
	r.log.Debug("Preparing statement to delete text secret from repository %s", secretID)
	query := "DELETE FROM login_password_secrets WHERE Id = $1"
	var args = []interface{}{secretID}
	secretType := "login-password"
	r.log.Debug("Preparing statement to delete login-password from Repository: %s ", args[0])
	return r.deleteSecret(ctx, secretType, query, args, tx)
}

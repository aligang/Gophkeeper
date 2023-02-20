package sql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aligang/Gophkeeper/pkg/common/secret/instance"
	"github.com/aligang/Gophkeeper/pkg/server/repository/repositoryerrors"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
)

func (r *Repository) AddTextSecret(ctx context.Context, s *instance.TextSecret, tx *transaction.DBTransaction) error {
	query := "INSERT INTO text_secrets (Id, AccountId, CreatedAt, ModifiedAt, Text) VALUES($1, $2, $3, $4, $5)"
	args := []any{s.Id, s.AccountId, s.CreatedAt, s.ModifiedAt, s.Text}
	secretType := "text"
	r.log.Debug("Preparing statement to create text secret to Repository: %+s", args[0])
	return r.addSecret(ctx, secretType, query, args, tx)
}

func (r *Repository) UpdateTextSecret(ctx context.Context, s *instance.TextSecret, tx *transaction.DBTransaction) error {
	query := "UPDATE text_secrets SET Id = $1, AccountId = $2, CreatedAt = $3, ModifiedAt = $4, Text = $5 WHERE Id = $6"
	args := []any{s.Id, s.AccountId, s.CreatedAt, s.ModifiedAt, s.Text, s.Id}
	secretType := "text"
	r.log.Debug("Preparing statement to update text secret to Repository: %s", args[0])
	return r.updateSecret(ctx, secretType, query, args, tx)
}

func (r *Repository) GetTextSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance.TextSecret, error) {
	query := "SELECT * FROM text_secrets WHERE id = $1"
	var args = []interface{}{secretID}
	secretType := "text"
	r.log.Debug("Preparing statement to fetch text from Repository: %s ", args[0])
	statement, err := r.prepareStatement(ctx, query, tx)
	s := &instance.TextSecret{}

	r.log.Debug("Executing statement to %s secret text from Repository: %s ", secretType, args[0])
	err = statement.GetContext(ctx, s, args...)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		r.log.Warn("Database response is empty")
		return nil, repositoryerrors.ErrNoContent
	case err != nil:
		r.log.Crit("Error during decoding database response: %s", err.Error())
		return nil, err
	default:
		return s, nil
	}
}

func (r *Repository) ListTextSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance.TextSecret, error) {
	query := "SELECT * FROM text_secrets WHERE accountId = $1"
	args := []interface{}{accountID}
	secretType := "text"
	r.log.Debug("Preparing statement to fetch login-password secrets from Repository: %s ", args[0])
	statement, err := r.prepareStatement(ctx, query, tx)

	var texts []*instance.TextSecret
	r.log.Debug("Executing statement to fetch %s secrets from Repository for account: %s", secretType, args[0])
	err = statement.SelectContext(ctx, &texts, args...)
	if err != nil {
		r.log.Warn("Error During statement Execution %s with %s", query, args[0])
		return texts, err
	}
	return texts, nil
}

func (r *Repository) DeleteTextSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error {
	r.log.Debug("Preparing statement to delete text secret from repository %s", secretID)
	query := "DELETE FROM text_secrets WHERE Id = $1"
	var args = []interface{}{secretID}
	secretType := "text"
	r.log.Debug("Preparing statement to delete text from Repository: %s ", args[0])
	return r.deleteSecret(ctx, secretType, query, args, tx)
}

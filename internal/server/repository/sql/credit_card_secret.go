package sql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aligang/Gophkeeper/internal/common/secret/instance"
	"github.com/aligang/Gophkeeper/internal/server/repository/repositoryerrors"
	"github.com/aligang/Gophkeeper/internal/server/repository/transaction"
)

func (r *Repository) AddCreditCardSecret(ctx context.Context, s *instance.CreditCardSecret, tx *transaction.DBTransaction) error {
	query := "INSERT INTO credit_card_secrets (Id, AccountId, CreatedAt, ModifiedAt, CardNumber, CardHolder, ValidTill, Cvc) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"
	args := []any{s.Id, s.AccountId, s.CreatedAt, s.ModifiedAt, s.CardNumber, s.CardHolder, s.ValidTill, s.Cvc}

	secretType := "credit-card"
	r.log.Debug("Preparing statement to create credit-card secret to Repository: %+s", args[0])
	return r.addSecret(ctx, secretType, query, args, tx)
}

func (r *Repository) UpdateCreditCardSecret(ctx context.Context, s *instance.CreditCardSecret, tx *transaction.DBTransaction) error {
	query := "UPDATE credit_card_secrets SET Id = $1, AccountId = $2, CreatedAt = $3, ModifiedAt = $4, CardNumber = $5, CardHolder = $6, ValidTill = $7, Cvc = $8 WHERE Id = $9"
	args := []any{s.Id, s.AccountId, s.CreatedAt, s.ModifiedAt, s.CardNumber, s.CardHolder, s.ValidTill, s.Cvc, s.Id}
	secretType := "credit-card"
	r.log.Debug("Preparing statement to update credit-card secret to Repository: %s", args[0])
	return r.updateSecret(ctx, secretType, query, args, tx)
}

func (r *Repository) GetCreditCardSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance.CreditCardSecret, error) {
	query := "SELECT * FROM credit_card_secrets WHERE id = $1"
	var args = []interface{}{secretID}
	r.log.Debug("Preparing statement to fetch login password from Repository: ")
	statement, err := r.prepareStatement(ctx, query, tx)
	s := &instance.CreditCardSecret{}

	r.log.Debug("Executing statement to login password secret  from Repository")
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

func (r *Repository) ListCreditCardSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance.CreditCardSecret, error) {
	r.log.Debug("Preparing statement to fetch account credit-card secrets from Repository")
	query := "SELECT * FROM credit_card_secrets WHERE accountId = $1"
	args := []interface{}{accountID}
	statement, err := r.prepareStatement(ctx, query, tx)

	var secrets []*instance.CreditCardSecret
	r.log.Debug("Executing statement to fetch credit-card secret from Repository")
	err = statement.SelectContext(ctx, &secrets, args...)
	if err != nil {
		r.log.Warn("Error During statement Execution %s with %s", query, args[0])
		return secrets, err
	}
	return secrets, nil
}

func (r *Repository) DeleteCreditCardSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error {
	r.log.Debug("Preparing statement to delete text secret from repository %s", secretID)
	query := "DELETE FROM credit_card_secrets WHERE Id = $1"
	var args = []interface{}{secretID}
	secretType := "credit-card"
	r.log.Debug("Preparing statement to delete credit-card from Repository: %s ", args[0])
	return r.deleteSecret(ctx, secretType, query, args, tx)
}

package inmemory

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/common/secret/instance"
	"github.com/aligang/Gophkeeper/internal/server/repository/repositoryerrors"
	"github.com/aligang/Gophkeeper/internal/server/repository/transaction"
)

type creditCardRecord struct {
	number     string
	cardholder string
	validTill  string
	cvc        string
}

func convertCreditCardSecretInstance(instance *instance.CreditCardSecret) *SecretRecord {
	return &SecretRecord{
		id:         instance.Id,
		accountId:  instance.AccountId,
		createdAt:  instance.CreatedAt,
		modifiedAt: instance.ModifiedAt,
		creditCardRecord: creditCardRecord{
			number:     instance.CardNumber,
			cardholder: instance.CardHolder,
			validTill:  instance.ValidTill,
			cvc:        instance.Cvc,
		},
	}
}

func convertCreditCardSecretRecord(r *SecretRecord) *instance.CreditCardSecret {
	return &instance.CreditCardSecret{
		BaseSecret: instance.BaseSecret{
			Id:         r.id,
			AccountId:  r.accountId,
			CreatedAt:  r.createdAt,
			ModifiedAt: r.modifiedAt,
		},
		CardNumber: r.creditCardRecord.number,
		CardHolder: r.creditCardRecord.cardholder,
		ValidTill:  r.creditCardRecord.validTill,
		Cvc:        r.creditCardRecord.cvc,
	}
}

func (r *Repository) AddCreditCardSecret(ctx context.Context, s *instance.CreditCardSecret, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("CreditCardSecret", "Add")
	logger.Debug("Adding new secret %s", s.Id)
	record := convertCreditCardSecretInstance(s)
	r.creditCardSecrets[record.id] = record
	if _, ok := r.accountCreditCardSecrets[record.accountId]; !ok {
		r.accountCreditCardSecrets[record.accountId] = map[string]any{}
	}
	r.accountCreditCardSecrets[record.accountId][s.Id] = nil
	logger.Debug("Secret %s successfully added", s.Id)
	return nil
}

func (r *Repository) UpdateCreditCardSecret(ctx context.Context, s *instance.CreditCardSecret, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("CreditCardSecret", "Update")
	logger.Debug("Updating secret %s", s.Id)
	r.creditCardSecrets[s.Id] = convertCreditCardSecretInstance(s)
	logger.Debug("Secret %s successfully updated", s.Id)
	return nil
}

func (r *Repository) ListCreditCardSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance.CreditCardSecret, error) {
	logger := r.log.GetSubLogger("CreditCardSecret", "List")
	logger.Debug("Listing secrets")
	secrets := []*instance.CreditCardSecret{}
	secretIDs, exists := r.accountCreditCardSecrets[accountID]
	if !exists {
		return secrets, nil
	}
	for ID, _ := range secretIDs {
		secrets = append(secrets, convertCreditCardSecretRecord(r.creditCardSecrets[ID]))
	}
	return secrets, nil
}

func (r *Repository) GetCreditCardSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance.CreditCardSecret, error) {
	logger := r.log.GetSubLogger("CreditCardSecret", "Get")
	logger.Debug("Fetching secret %s", secretID)
	record, ok := r.creditCardSecrets[secretID]
	if !ok {
		r.log.Debug("secret %s not found", secretID)
		return nil, repositoryerrors.ErrRecordNotFound
	}
	logger.Debug("Secret %s successfully fetched", secretID)
	return convertCreditCardSecretRecord(record), nil
}

func (r *Repository) DeleteCreditCardSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("CreditCardSecret", "Delete")
	logger.Debug("Deleting secret %s", secretID)
	accountID := r.creditCardSecrets[secretID].accountId
	delete(r.accountCreditCardSecrets[accountID], secretID)
	if len(r.accountCreditCardSecrets[accountID]) == 0 {
		delete(r.accountCreditCardSecrets, accountID)
	}
	delete(r.creditCardSecrets, secretID)
	logger.Debug("Secret %s successfully Deleted", secretID)
	return nil
}

package inmemory

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/secret/instance"
	"github.com/aligang/Gophkeeper/pkg/server/repository/repositoryerrors"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
)

type loginPasswordRecord struct {
	login    string
	password string
}

func convertLoginPasswordSecretInstance(instance *instance.LoginPasswordSecret) *SecretRecord {
	return &SecretRecord{
		id:         instance.Id,
		accountId:  instance.AccountId,
		createdAt:  instance.CreatedAt,
		modifiedAt: instance.ModifiedAt,
		loginPasswordRecord: loginPasswordRecord{
			login:    instance.Login,
			password: instance.Password,
		},
	}
}

func convertLoginPasswordSecretRecord(r *SecretRecord) *instance.LoginPasswordSecret {
	return &instance.LoginPasswordSecret{
		BaseSecret: instance.BaseSecret{
			Id:         r.id,
			AccountId:  r.accountId,
			CreatedAt:  r.createdAt,
			ModifiedAt: r.modifiedAt,
		},
		Login:    r.loginPasswordRecord.login,
		Password: r.loginPasswordRecord.password,
	}
}

func (r *Repository) AddLoginPasswordSecret(ctx context.Context, s *instance.LoginPasswordSecret, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("LoginPasswordSecret", "Add")
	logger.Debug("Adding new secret %s", s.Id)
	record := convertLoginPasswordSecretInstance(s)
	r.loginPasswordSecrets[record.id] = record
	if _, ok := r.accountLoginPasswordSecrets[record.accountId]; !ok {
		r.accountLoginPasswordSecrets[record.accountId] = map[string]any{}
	}
	r.accountLoginPasswordSecrets[record.accountId][s.Id] = nil
	logger.Debug("Secret %s successfully added", s.Id)
	return nil
}

func (r *Repository) UpdateLoginPasswordSecret(ctx context.Context, s *instance.LoginPasswordSecret, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("LoginPasswordSecret", "Update")
	logger.Debug("Updating secret %s", s.Id)
	r.loginPasswordSecrets[s.Id] = convertLoginPasswordSecretInstance(s)
	logger.Debug("Secret %s successfully updated", s.Id)
	return nil
}

func (r *Repository) ListLoginPasswordSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance.LoginPasswordSecret, error) {
	logger := r.log.GetSubLogger("LoginPasswordSecret", "List")
	logger.Debug("Listing secrets")
	secrets := []*instance.LoginPasswordSecret{}
	secretIDs, exists := r.accountLoginPasswordSecrets[accountID]
	if !exists {
		return secrets, nil
	}
	for ID, _ := range secretIDs {
		secrets = append(secrets, convertLoginPasswordSecretRecord(r.loginPasswordSecrets[ID]))
	}
	return secrets, nil
}

func (r *Repository) GetLoginPasswordSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance.LoginPasswordSecret, error) {
	logger := r.log.GetSubLogger("LoginPasswordSecret", "Get")
	logger.Debug("Fetching secret %s", secretID)
	record, ok := r.loginPasswordSecrets[secretID]
	if !ok {
		r.log.Debug("secret %s not found", secretID)
		return nil, repositoryerrors.ErrRecordNotFound
	}
	logger.Debug("Secret %s successfully fetched", secretID)
	return convertLoginPasswordSecretRecord(record), nil
}

func (r *Repository) DeleteLoginPasswordSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("LoginPasswordSecret", "Delete")
	logger.Debug("Deleting secret %s", secretID)
	accountID := r.loginPasswordSecrets[secretID].accountId
	delete(r.accountLoginPasswordSecrets[accountID], secretID)
	if len(r.accountLoginPasswordSecrets[accountID]) == 0 {
		delete(r.accountLoginPasswordSecrets, accountID)
	}
	delete(r.loginPasswordSecrets, secretID)
	logger.Debug("Secret %s successfully Deleted", secretID)
	return nil
}

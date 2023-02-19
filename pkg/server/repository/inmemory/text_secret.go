package inmemory

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/common/secret/instance"
	"github.com/aligang/Gophkeeper/pkg/server/repository/repositoryerrors"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
)

func convertTextSecretInstance(r *instance.TextSecret) *SecretRecord {
	return &SecretRecord{
		id:         r.Id,
		accountId:  r.AccountId,
		createdAt:  r.CreatedAt,
		modifiedAt: r.ModifiedAt,
		text:       r.Text,
	}
}

func convertTextSecretRecord(r *SecretRecord) *instance.TextSecret {
	return &instance.TextSecret{
		BaseSecret: instance.BaseSecret{
			Id:         r.id,
			AccountId:  r.accountId,
			CreatedAt:  r.createdAt,
			ModifiedAt: r.modifiedAt,
		},
		Text: r.text,
	}
}

func (r *Repository) AddTextSecret(ctx context.Context, s *instance.TextSecret, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("TextSecret", "Add")
	logger.Debug("Adding new secret %s", s.Id)
	record := convertTextSecretInstance(s)
	r.textSecrets[record.id] = record
	if _, ok := r.accountTextSecrets[record.accountId]; !ok {
		r.accountTextSecrets[record.accountId] = map[string]any{}
	}
	r.accountTextSecrets[record.accountId][s.Id] = nil
	logger.Debug("Secret %s successfully added", s.Id)
	return nil
}

func (r *Repository) UpdateTextSecret(ctx context.Context, s *instance.TextSecret, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("TextSecret", "Update")
	logger.Debug("Updating secret %s", s.Id)
	r.textSecrets[s.Id] = convertTextSecretInstance(s)
	logger.Debug("Secret %s successfully updated", s.Id)
	return nil
}

func (r *Repository) ListTextSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance.TextSecret, error) {
	logger := r.log.GetSubLogger("TextSecret", "List")
	logger.Debug("Listing secrets")
	secrets := []*instance.TextSecret{}
	secretIDs, exists := r.accountTextSecrets[accountID]
	if !exists {
		return secrets, nil
	}
	for ID, _ := range secretIDs {
		secrets = append(secrets, convertTextSecretRecord(r.textSecrets[ID]))
	}
	return secrets, nil
}

func (r *Repository) GetTextSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance.TextSecret, error) {
	logger := r.log.GetSubLogger("TextSecret", "Get")
	logger.Debug("Fetching secret %s", secretID)
	record, ok := r.textSecrets[secretID]
	if !ok {
		r.log.Debug("secret %s not found", secretID)
		return nil, repositoryerrors.ErrRecordNotFound
	}
	logger.Debug("Secret %s successfully fetched", secretID)
	return convertTextSecretRecord(record), nil
}

func (r *Repository) DeleteTextSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("TextSecret", "Delete")
	logger.Debug("Deleting secret %s", secretID)
	accountID := r.textSecrets[secretID].accountId
	delete(r.accountTextSecrets[accountID], secretID)
	if len(r.accountTextSecrets[accountID]) == 0 {
		delete(r.accountTextSecrets, accountID)
	}
	delete(r.textSecrets, secretID)
	logger.Debug("Secret %s successfully Deleted", secretID)
	return nil
}

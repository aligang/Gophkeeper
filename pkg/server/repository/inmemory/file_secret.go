package inmemory

import (
	"context"
	instance2 "github.com/aligang/Gophkeeper/pkg/common/secret/instance"
	"github.com/aligang/Gophkeeper/pkg/server/repository/repositoryerrors"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
	"time"
)

type fileRecord struct {
	objectId string
}

func convertFileSecretInstance(r *instance2.FileSecret) *SecretRecord {
	return &SecretRecord{
		id:         r.Id,
		accountId:  r.AccountId,
		createdAt:  r.CreatedAt,
		modifiedAt: r.ModifiedAt,
		fileRecord: fileRecord{objectId: r.ObjectId},
	}
}

func convertFileSecretRecord(r *SecretRecord) *instance2.FileSecret {
	return &instance2.FileSecret{
		BaseSecret: instance2.BaseSecret{
			Id:         r.id,
			AccountId:  r.accountId,
			CreatedAt:  r.createdAt,
			ModifiedAt: r.modifiedAt,
		},
		ObjectId: r.fileRecord.objectId,
	}
}

func (r *Repository) AddFileSecret(ctx context.Context, s *instance2.FileSecret, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("FileSecret", "Add")
	logger.Debug("Adding new secret %s", s.Id)
	r.fileSecrets[s.Id] = convertFileSecretInstance(s)
	if _, ok := r.accountFileSecrets[s.AccountId]; !ok {
		r.accountFileSecrets[s.AccountId] = map[string]any{}
	}
	r.accountFileSecrets[s.AccountId][s.Id] = nil
	logger.Debug("Secret %s is successfully added", s.Id)
	return nil
}

func (r *Repository) UpdateFileSecret(ctx context.Context, s *instance2.FileSecret, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("FileSecret", "Add")
	logger.Debug("Updating existing secret %s", s.Id)
	r.fileSecrets[s.Id] = convertFileSecretInstance(s)
	logger.Debug("Secret %s is successfully updated", s.Id)
	return nil
}

func (r *Repository) ListFileSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance2.FileSecret, error) {
	logger := r.log.GetSubLogger("FileSecret", "List")
	logger.Debug("Listing secrets")
	secrets := []*instance2.FileSecret{}
	secretIDs, exists := r.accountFileSecrets[accountID]
	if !exists {
		logger.Debug("No records were found")
		return secrets, nil
	}
	for ID, _ := range secretIDs {
		secrets = append(secrets, convertFileSecretRecord(r.fileSecrets[ID]))
	}
	logger.Debug("Secrets successfully listed")
	return secrets, nil
}

func (r *Repository) GetFileSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance2.FileSecret, error) {
	logger := r.log.GetSubLogger("FileSecret", "Get")
	logger.Debug("Fetching existing secret %s", secretID)
	s, ok := r.fileSecrets[secretID]
	if !ok {
		r.log.Debug("secret %s not found", secretID)
		return nil, repositoryerrors.ErrRecordNotFound
	}
	logger.Debug("Secret %s is successfully fetched", secretID)
	return convertFileSecretRecord(s), nil
}

func (r *Repository) DeleteFileSecretFromDeletionQueue(ctx context.Context, secretID string, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("FileSecret", "DeleteObjectRecordFromDeletionQueue")
	logger.Debug("Deleting file secret from deletion queue %s ", secretID)
	delete(r.fileDeletionQueue, secretID)
	logger.Debug("Deletion of %s from deletion queue Succeed", secretID)
	return nil
}

func (r *Repository) MoveFileSecretToDeletionQueue(ctx context.Context, objectId string, ts time.Time, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("FileSecret", "MoveObjectToDeletionQueue")
	logger.Debug("Moving object to deletion queue: %s", objectId)
	r.fileDeletionQueue[objectId] = ts
	logger.Debug("Movement of  object %s from deletion queue Succeed", objectId)
	return nil
}

func (r *Repository) ListFileDeletionQueue(ctx context.Context, tx *transaction.DBTransaction) ([]*instance2.DeletionQueueElement, error) {
	logger := r.log.GetSubLogger("FileSecret", "ListDeletionQueue")
	logger.Debug("Listing elements of deletion queue")
	q := []*instance2.DeletionQueueElement{}
	for id, ts := range r.fileDeletionQueue {
		logger.Debug("Found %s ", id)
		q = append(q, &instance2.DeletionQueueElement{ObjectId: id, DeletedAt: ts})
	}
	logger.Debug("Content of file deletion queue is successfully fetched: %d elements", len(q))
	return q, nil
}

func (r *Repository) DeleteFileSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("FileSecret", "Delete")
	logger.Debug("Deleting secret %s", secretID)
	accountID := r.fileSecrets[secretID].accountId
	delete(r.accountFileSecrets[accountID], secretID)
	if len(r.accountFileSecrets[accountID]) == 0 {
		delete(r.accountFileSecrets, accountID)
	}
	delete(r.fileSecrets, secretID)
	logger.Debug("Secret record %s is successfully deleted", secretID)
	return nil
}

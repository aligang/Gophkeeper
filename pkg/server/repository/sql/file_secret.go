package sql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aligang/Gophkeeper/pkg/common/secret/instance"
	"github.com/aligang/Gophkeeper/pkg/server/repository/repositoryerrors"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
	"time"
)

func (r *Repository) AddFileSecret(ctx context.Context, s *instance.FileSecret, tx *transaction.DBTransaction) error {
	query := "INSERT INTO file_secrets (Id, AccountId, CreatedAt, ModifiedAt, ObjectId) VALUES($1, $2, $3, $4, $5)"
	args := []any{s.Id, s.AccountId, s.CreatedAt, s.ModifiedAt, s.ObjectId}

	secretType := "file"
	r.log.Debug("Preparing statement to create file secret to Repository: %+s", args[0])
	return r.addSecret(ctx, secretType, query, args, tx)
}

func (r *Repository) UpdateFileSecret(ctx context.Context, s *instance.FileSecret, tx *transaction.DBTransaction) error {
	query := "UPDATE file_secrets SET Id = $1, AccountId = $2, CreatedAt = $3, ModifiedAt = $4, ObjectId = $5 WHERE Id = $6"
	args := []any{s.Id, s.AccountId, s.CreatedAt, s.ModifiedAt, s.ObjectId, s.Id}
	secretType := "file"
	r.log.Debug("Preparing statement to update file secret to Repository: %s", args[0])
	return r.updateSecret(ctx, secretType, query, args, tx)
}

func (r *Repository) GetFileSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance.FileSecret, error) {
	query := "SELECT * FROM file_secrets WHERE id = $1"
	var args = []interface{}{secretID}
	r.log.Debug("Preparing statement to fetch login password from Repository: ")
	statement, err := r.prepareStatement(ctx, query, tx)
	s := &instance.FileSecret{}

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

func (r *Repository) ListFileSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance.FileSecret, error) {
	r.log.Debug("Preparing statement to fetch account file secrets from Repository")
	query := "SELECT * FROM file_secrets WHERE accountId = $1"
	args := []interface{}{accountID}
	statement, err := r.prepareStatement(ctx, query, tx)

	var secrets []*instance.FileSecret
	r.log.Debug("Executing statement to fetch account file secrets from Repository")
	err = statement.SelectContext(ctx, &secrets, args...)
	if err != nil {
		r.log.Warn("Error During statement Execution %s with %s", query, args[0])
		return secrets, err
	}
	return secrets, nil
}

func (r *Repository) DeleteFileSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error {
	r.log.Debug("Preparing statement to delete file secret from repository %s", secretID)
	query := "DELETE FROM file_secrets WHERE Id = $1"
	var args = []interface{}{secretID}
	secretType := "file"
	r.log.Debug("Preparing statement to delete file from Repository: %s ", args[0])
	return r.deleteSecret(ctx, secretType, query, args, tx)
}

func (r *Repository) MoveFileSecretToDeletionQueue(ctx context.Context, objectId string, ts time.Time, tx *transaction.DBTransaction) error {
	r.log.Debug("Executing statement to move file to file deletion queue: %s", objectId)
	query := "INSERT INTO file_deletion_queue (ObjectId, DeletedAt) VALUES($1, $2)"
	args := []any{objectId, ts}
	statement, err := r.prepareStatement(ctx, query, tx)

	_, err = statement.ExecContext(ctx, args...)
	if err != nil {
		r.log.Crit("Error During statement Execution of movement file to deletion queue %s secret with id %s: %s", args[0], err.Error())
		return err
	}
	return nil
}

func (r *Repository) ListFileDeletionQueue(ctx context.Context, tx *transaction.DBTransaction) ([]*instance.DeletionQueueElement, error) {
	r.log.Debug("Preparing statement to fetch account file secrets from Repository")
	query := "SELECT ObjectId, DeletedAt FROM file_deletion_queue"
	args := []interface{}{}
	statement, err := r.prepareStatement(ctx, query, tx)

	var filesToDelete []*instance.DeletionQueueElement
	r.log.Debug("Executing statement to fetch file secret from Repository")
	err = statement.SelectContext(ctx, &filesToDelete, args...)
	if err != nil {
		r.log.Crit("Error During statement Execution %s with %s", query)
		return filesToDelete, err
	}
	return filesToDelete, nil
}

func (r *Repository) DeleteFileSecretFromDeletionQueue(ctx context.Context, objectId string, tx *transaction.DBTransaction) error {
	r.log.Debug("Preparing statement to delete file secret from deletion queue %s", objectId)
	query := "DELETE FROM file_deletion_queue WHERE ObjectId = $1"
	var args = []interface{}{objectId}
	statement, err := r.prepareStatement(ctx, query, tx)
	_, err = statement.ExecContext(ctx, args...)
	if err != nil {
		r.log.Crit("Error During statement Execution %s with %s", query, args[0])
		return err
	}
	return nil
}

package sql

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/secret/instance"
	"time"
)

func (r *Repository) AddFileSecret(ctx context.Context, s *instance.FileSecret) error {
	return nil
}

func (r *Repository) UpdateFileSecret(ctx context.Context, s *instance.FileSecret) error {
	return nil
}

func (r *Repository) ListFileSecrets(ctx context.Context, accountID string) ([]*instance.FileSecret, error) {
	return nil, nil
}

func (r *Repository) GetFileSecret(ctx context.Context, secretID string) (*instance.FileSecret, error) {
	return nil, nil
}

func (r *Repository) MoveFileSecretToDeletionQueue(ctx context.Context, objectId string, ts time.Time) error {
	return nil
}

func (r *Repository) ListFileDeletionQueue(ctx context.Context) ([]*instance.DeletionQueueElement, error) {
	return nil, nil
}

func (r *Repository) DeleteFileSecret(ctx context.Context, secretID string) error {
	return nil
}

func (r *Repository) DeleteFileSecretFromDeletionQueue(ctx context.Context, secretID string) error {
	return nil
}

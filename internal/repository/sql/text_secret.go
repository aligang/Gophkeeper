package sql

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/secret/instance"
)

func (r *Repository) AddTextSecret(ctx context.Context, s *instance.TextSecret) error {
	return nil
}

func (r *Repository) UpdateTextSecret(ctx context.Context, s *instance.TextSecret) error {
	return nil
}

func (r *Repository) GetTextSecret(ctx context.Context, secretID string) (*instance.TextSecret, error) {
	return nil, nil
}

func (r *Repository) ListTextSecrets(ctx context.Context, accountID string) ([]*instance.TextSecret, error) {
	return nil, nil
}

func (r *Repository) DeleteTextSecret(ctx context.Context, secretID string) error {
	return nil
}

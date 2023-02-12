package sql

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/secret/instance"
)

func (r *Repository) AddLoginPasswordSecret(ctx context.Context, s *instance.LoginPasswordSecret) error {
	return nil
}

func (r *Repository) UpdateLoginPasswordSecret(ctx context.Context, s *instance.LoginPasswordSecret) error {
	return nil
}

func (r *Repository) GetLoginPasswordSecret(ctx context.Context, secretID string) (*instance.LoginPasswordSecret, error) {
	return nil, nil
}

func (r *Repository) ListLoginPasswordSecrets(ctx context.Context, accountID string) ([]*instance.LoginPasswordSecret, error) {
	return nil, nil
}

func (r *Repository) DeleteLoginPasswordSecret(ctx context.Context, secretID string) error {
	return nil
}

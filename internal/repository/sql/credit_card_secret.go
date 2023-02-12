package sql

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/secret/instance"
)

func (r *Repository) AddCreditCardSecret(ctx context.Context, s *instance.CreditCardSecret) error {
	return nil
}

func (r *Repository) UpdateCreditCardSecret(ctx context.Context, s *instance.CreditCardSecret) error {
	return nil
}

func (r *Repository) GetCreditCardSecret(ctx context.Context, secretID string) (*instance.CreditCardSecret, error) {
	return nil, nil
}

func (r *Repository) ListCreditCardSecrets(ctx context.Context, accountID string) ([]*instance.CreditCardSecret, error) {
	return nil, nil
}

func (r *Repository) DeleteCreditCardSecret(ctx context.Context, secretID string) error {
	return nil
}

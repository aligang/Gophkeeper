package sql

import (
	"context"
	accountInstance "github.com/aligang/Gophkeeper/internal/account/instance"
	tokenInstance "github.com/aligang/Gophkeeper/internal/token/instance"
)

func (r *Repository) Register(ctx context.Context, account *accountInstance.Account) error {
	return nil
}

func (r *Repository) GetAccountByLogin(ctx context.Context, login string) (*accountInstance.Account, error) {
	return nil, nil
}

func (r *Repository) GetAccountById(ctx context.Context, accountID string) (*accountInstance.Account, error) {
	return nil, nil
}

func (r *Repository) AddToken(ctx context.Context, t *tokenInstance.Token) error {
	return nil
}

func (r *Repository) GetToken(ctx context.Context, tokenValue string) (*tokenInstance.Token, error) {
	return nil, nil
}

func (r *Repository) ListAccountTokens(ctx context.Context, accountID string) ([]*tokenInstance.Token, error) {
	return nil, nil
}

func (r *Repository) ListTokens(ctx context.Context) ([]*tokenInstance.Token, error) {
	return nil, nil
}

func (r *Repository) DeleteToken(ctx context.Context, t *tokenInstance.Token) error {
	return nil
}

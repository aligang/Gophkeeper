package inmemory

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/repository/repositoryerrors"
	"github.com/aligang/Gophkeeper/internal/repository/transaction"
	tokenInstance "github.com/aligang/Gophkeeper/internal/token/instance"
	"sort"
	"time"
)

type tokenRecord struct {
	id       string
	issuedAt time.Time
	value    string
	owner    string
}

func convertTokenInstance(i *tokenInstance.Token) *tokenRecord {
	return &tokenRecord{
		id:       i.Id,
		value:    i.TokenValue,
		owner:    i.Owner,
		issuedAt: i.IssuedAt,
	}
}

func convertTokenRecord(r *tokenRecord) *tokenInstance.Token {
	return &tokenInstance.Token{
		Id:         r.id,
		TokenValue: r.value,
		Owner:      r.owner,
		IssuedAt:   r.issuedAt,
	}
}

func (r *Repository) GetToken(ctx context.Context, tokenValue string, tx *transaction.DBTransaction) (*tokenInstance.Token, error) {
	logger := r.log.GetSubLogger("Token", "Get")
	logger.Debug("Fetching token record")
	tokenRecord, ok := r.tokens[tokenValue]
	if !ok {
		logger.Debug("Could not find token")
		return nil, repositoryerrors.ErrRecordNotFound
	}
	logger.Debug("Token %s successfully fetched", tokenRecord.id)
	return convertTokenRecord(tokenRecord), nil
}

func (r *Repository) AddToken(ctx context.Context, instance *tokenInstance.Token, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("Token", "Add")
	logger.Debug("Adding token record %s", instance.Id)
	r.tokens[instance.TokenValue] = convertTokenInstance(instance)
	_, ok := r.accountTokens[instance.Owner]
	if !ok {
		r.accountTokens[instance.Owner] = map[string]interface{}{}
	}
	logger.Debug("Token record %s added to databes", instance.Id)
	r.accountTokens[instance.Owner][instance.TokenValue] = nil
	return nil
}

func (r *Repository) ListAccountTokens(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*tokenInstance.Token, error) {
	logger := r.log.GetSubLogger("AccountToken", "List")
	logger.Debug("Listing record")
	tokens, ok := r.accountTokens[accountID]
	tokenInstances := []*tokenInstance.Token{}
	if !ok {
		return tokenInstances, nil
	}

	for t, _ := range tokens {
		tokenRecord, ok := r.tokens[t]
		if !ok {
			logger.Debug("token is not found withing records, database is inconsistent state")
		}
		tokenInstances = append(tokenInstances, convertTokenRecord(tokenRecord))
	}

	sort.Slice(tokenInstances, func(i, j int) bool {
		return tokenInstances[i].IssuedAt.After(tokenInstances[i].IssuedAt)
	})
	logger.Debug("Token records of account %s successfully listed", accountID)
	return tokenInstances, nil
}

func (r *Repository) ListTokens(ctx context.Context, tx *transaction.DBTransaction) ([]*tokenInstance.Token, error) {
	logger := r.log.GetSubLogger("Token", "List")
	logger.Debug("Listing record")

	tokenInstances := []*tokenInstance.Token{}
	for _, tokenRecord := range r.tokens {
		tokenInstances = append(tokenInstances, convertTokenRecord(tokenRecord))
	}

	logger.Debug("Token records successfully listed")
	return tokenInstances, nil
}

func (r *Repository) DeleteToken(ctx context.Context, t *tokenInstance.Token, tx *transaction.DBTransaction) error {
	logger := r.log.GetSubLogger("Token", "Delete")
	logger.Debug("Deleting token %s", t.Id)
	delete(r.tokens, t.TokenValue)

	delete(r.accountTokens[t.Owner], t.TokenValue)
	if len(r.accountTokens[t.Owner]) == 0 {
		delete(r.accountTokens, t.Owner)
	}
	logger.Debug("Token %s successfully deleted", t.Id)
	return nil
}

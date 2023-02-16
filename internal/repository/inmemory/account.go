package inmemory

import (
	"context"
	"errors"
	accountInstance "github.com/aligang/Gophkeeper/internal/account/instance"
	"github.com/aligang/Gophkeeper/internal/repository/transaction"
	"time"
)

type accountRecord struct {
	id            string
	createdAt     time.Time
	login         string
	password      string
	encryptionKey string
}

func convertAccountInstance(a *accountInstance.Account) *accountRecord {
	return &accountRecord{
		id:            a.Id,
		password:      a.Password,
		login:         a.Login,
		createdAt:     a.CreatedAt,
		encryptionKey: a.EncryptionKey,
	}
}

func convertAccountRecord(r *accountRecord) *accountInstance.Account {
	return &accountInstance.Account{
		Id:            r.id,
		Password:      r.password,
		Login:         r.login,
		CreatedAt:     r.createdAt,
		EncryptionKey: r.encryptionKey,
	}
}

func (r *Repository) Register(ctx context.Context, instance *accountInstance.Account, tx *transaction.DBTransaction) error {
	r.log.Debug("Registering New Account")
	r.log.Debug("Account id: %s", instance.Id)
	r.accounts[instance.Id] = convertAccountInstance(instance)
	r.log.Debug("Account id: %s", instance.Login)
	r.loginToIdMapping[instance.Login] = instance.Id
	r.log.Debug("Finished Registering New Account")
	return nil
}

func (r *Repository) GetAccountByLogin(ctx context.Context, login string, tx *transaction.DBTransaction) (*accountInstance.Account, error) {
	r.log.Debug("Getting Account using Login: %s", login)
	id, exists := r.loginToIdMapping[login]
	if !exists {
		r.log.Debug("Account does not exists")
		return nil, errors.New("account does not exists")
	}

	return r.GetAccountById(ctx, id, tx)
}

func (r *Repository) GetAccountById(ctx context.Context, accountID string, tx *transaction.DBTransaction) (*accountInstance.Account, error) {
	r.log.Debug("Get Account by Id from filerepository: %s", accountID)
	accountRecord, exists := r.accounts[accountID]
	if !exists {
		r.log.Debug("Account does not exists")
		return nil, errors.New("account does not exists")
	}
	r.log.Debug("Account record is found")
	return convertAccountRecord(accountRecord), nil
}

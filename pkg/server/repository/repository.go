package repository

import (
	"context"
	accountInstance "github.com/aligang/Gophkeeper/pkg/account/instance"
	"github.com/aligang/Gophkeeper/pkg/config"
	"github.com/aligang/Gophkeeper/pkg/logging"
	secretInstance "github.com/aligang/Gophkeeper/pkg/secret/instance"
	"github.com/aligang/Gophkeeper/pkg/server/repository/inmemory"
	"github.com/aligang/Gophkeeper/pkg/server/repository/sql"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
	tokenInstance "github.com/aligang/Gophkeeper/pkg/token/instance"
	"time"
)

type Storage interface {
	WithinTransaction(ctx context.Context, fn func(context.Context, *transaction.DBTransaction) error) error

	Register(ctx context.Context, account *accountInstance.Account, tx *transaction.DBTransaction) error
	GetAccountByLogin(ctx context.Context, login string, tx *transaction.DBTransaction) (*accountInstance.Account, error)
	GetAccountById(ctx context.Context, accountID string, tx *transaction.DBTransaction) (*accountInstance.Account, error)

	GetToken(ctx context.Context, tokenValue string, tx *transaction.DBTransaction) (*tokenInstance.Token, error)
	AddToken(ctx context.Context, t *tokenInstance.Token, tx *transaction.DBTransaction) error
	ListAccountTokens(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*tokenInstance.Token, error)
	ListTokens(ctx context.Context, tx *transaction.DBTransaction) ([]*tokenInstance.Token, error)
	DeleteToken(ctx context.Context, t *tokenInstance.Token, tx *transaction.DBTransaction) error

	AddTextSecret(ctx context.Context, s *secretInstance.TextSecret, tx *transaction.DBTransaction) error
	UpdateTextSecret(ctx context.Context, s *secretInstance.TextSecret, tx *transaction.DBTransaction) error
	GetTextSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*secretInstance.TextSecret, error)
	ListTextSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*secretInstance.TextSecret, error)
	DeleteTextSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error

	AddLoginPasswordSecret(ctx context.Context, s *secretInstance.LoginPasswordSecret, tx *transaction.DBTransaction) error
	UpdateLoginPasswordSecret(ctx context.Context, s *secretInstance.LoginPasswordSecret, tx *transaction.DBTransaction) error
	GetLoginPasswordSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*secretInstance.LoginPasswordSecret, error)
	ListLoginPasswordSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*secretInstance.LoginPasswordSecret, error)
	DeleteLoginPasswordSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error

	AddCreditCardSecret(ctx context.Context, s *secretInstance.CreditCardSecret, tx *transaction.DBTransaction) error
	UpdateCreditCardSecret(ctx context.Context, s *secretInstance.CreditCardSecret, tx *transaction.DBTransaction) error
	GetCreditCardSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*secretInstance.CreditCardSecret, error)
	ListCreditCardSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*secretInstance.CreditCardSecret, error)
	DeleteCreditCardSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error

	AddFileSecret(ctx context.Context, s *secretInstance.FileSecret, tx *transaction.DBTransaction) error
	UpdateFileSecret(ctx context.Context, s *secretInstance.FileSecret, tx *transaction.DBTransaction) error
	GetFileSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*secretInstance.FileSecret, error)
	ListFileSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*secretInstance.FileSecret, error)
	DeleteFileSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error

	MoveFileSecretToDeletionQueue(ctx context.Context, objectId string, ts time.Time, tx *transaction.DBTransaction) error
	ListFileDeletionQueue(ctx context.Context, tx *transaction.DBTransaction) ([]*secretInstance.DeletionQueueElement, error)
	DeleteFileSecretFromDeletionQueue(ctx context.Context, secretID string, tx *transaction.DBTransaction) error
}

func New(serverConfig *config.ServerConfig) Storage {
	var storage Storage
	logging.Debug("Initialization Storage")
	if serverConfig.GetRepositoryType() == config.RepositoryType_IN_MEMORY {
		storage = inmemory.New()
	} else if serverConfig.GetRepositoryType() == config.RepositoryType_SQL {
		storage = sql.New(serverConfig)
	}
	logging.Debug("Storage Initialization finished")
	return storage
}

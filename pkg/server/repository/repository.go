package repository

import (
	"context"
	accountInstance "github.com/aligang/Gophkeeper/pkg/common/account/instance"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	"github.com/aligang/Gophkeeper/pkg/common/secret/instance"
	tokenInstance "github.com/aligang/Gophkeeper/pkg/common/token/instance"
	"github.com/aligang/Gophkeeper/pkg/server/config"
	"github.com/aligang/Gophkeeper/pkg/server/repository/inmemory"
	"github.com/aligang/Gophkeeper/pkg/server/repository/sql"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
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

	AddTextSecret(ctx context.Context, s *instance.TextSecret, tx *transaction.DBTransaction) error
	UpdateTextSecret(ctx context.Context, s *instance.TextSecret, tx *transaction.DBTransaction) error
	GetTextSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance.TextSecret, error)
	ListTextSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance.TextSecret, error)
	DeleteTextSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error

	AddLoginPasswordSecret(ctx context.Context, s *instance.LoginPasswordSecret, tx *transaction.DBTransaction) error
	UpdateLoginPasswordSecret(ctx context.Context, s *instance.LoginPasswordSecret, tx *transaction.DBTransaction) error
	GetLoginPasswordSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance.LoginPasswordSecret, error)
	ListLoginPasswordSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance.LoginPasswordSecret, error)
	DeleteLoginPasswordSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error

	AddCreditCardSecret(ctx context.Context, s *instance.CreditCardSecret, tx *transaction.DBTransaction) error
	UpdateCreditCardSecret(ctx context.Context, s *instance.CreditCardSecret, tx *transaction.DBTransaction) error
	GetCreditCardSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance.CreditCardSecret, error)
	ListCreditCardSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance.CreditCardSecret, error)
	DeleteCreditCardSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error

	AddFileSecret(ctx context.Context, s *instance.FileSecret, tx *transaction.DBTransaction) error
	UpdateFileSecret(ctx context.Context, s *instance.FileSecret, tx *transaction.DBTransaction) error
	GetFileSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) (*instance.FileSecret, error)
	ListFileSecrets(ctx context.Context, accountID string, tx *transaction.DBTransaction) ([]*instance.FileSecret, error)
	DeleteFileSecret(ctx context.Context, secretID string, tx *transaction.DBTransaction) error

	MoveFileSecretToDeletionQueue(ctx context.Context, objectId string, ts time.Time, tx *transaction.DBTransaction) error
	ListFileDeletionQueue(ctx context.Context, tx *transaction.DBTransaction) ([]*instance.DeletionQueueElement, error)
	DeleteFileSecretFromDeletionQueue(ctx context.Context, secretID string, tx *transaction.DBTransaction) error
}

func New(serverConfig *config.Config) Storage {
	var storage Storage
	logging.Info("Initialization Storage")
	if serverConfig.GetRepositoryType() == config.RepositoryType_IN_MEMORY {
		storage = inmemory.New()
	} else if serverConfig.GetRepositoryType() == config.RepositoryType_SQL {
		storage = sql.New(serverConfig)
	}
	logging.Info("Storage Initialization finished")
	return storage
}

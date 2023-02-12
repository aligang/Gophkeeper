package repository

import (
	"context"
	accountInstance "github.com/aligang/Gophkeeper/internal/account/instance"
	"github.com/aligang/Gophkeeper/internal/config"
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/aligang/Gophkeeper/internal/repository/inmemory"
	"github.com/aligang/Gophkeeper/internal/repository/sql"
	"github.com/aligang/Gophkeeper/internal/repository/transaction"
	secretInstance "github.com/aligang/Gophkeeper/internal/secret/instance"
	tokenInstance "github.com/aligang/Gophkeeper/internal/token/instance"
	"time"
)

type Storage interface {
	WithinTransaction(ctx context.Context, fn func(context.Context, *transaction.DBTransaction) error) error

	Register(ctx context.Context, account *accountInstance.Account) error
	GetAccountByLogin(ctx context.Context, login string) (*accountInstance.Account, error)
	GetAccountById(ctx context.Context, accountID string) (*accountInstance.Account, error)

	GetToken(ctx context.Context, tokenValue string) (*tokenInstance.Token, error)
	AddToken(ctx context.Context, t *tokenInstance.Token) error
	ListAccountTokens(ctx context.Context, accountID string) ([]*tokenInstance.Token, error)
	ListTokens(ctx context.Context) ([]*tokenInstance.Token, error)
	DeleteToken(context.Context, *tokenInstance.Token) error

	AddTextSecret(ctx context.Context, s *secretInstance.TextSecret) error
	UpdateTextSecret(ctx context.Context, s *secretInstance.TextSecret) error
	GetTextSecret(ctx context.Context, secretID string) (*secretInstance.TextSecret, error)
	ListTextSecrets(ctx context.Context, accountID string) ([]*secretInstance.TextSecret, error)
	DeleteTextSecret(ctx context.Context, secretID string) error

	AddLoginPasswordSecret(ctx context.Context, s *secretInstance.LoginPasswordSecret) error
	UpdateLoginPasswordSecret(ctx context.Context, s *secretInstance.LoginPasswordSecret) error
	GetLoginPasswordSecret(ctx context.Context, secretID string) (*secretInstance.LoginPasswordSecret, error)
	ListLoginPasswordSecrets(ctx context.Context, accountID string) ([]*secretInstance.LoginPasswordSecret, error)
	DeleteLoginPasswordSecret(ctx context.Context, secretID string) error

	AddCreditCardSecret(ctx context.Context, s *secretInstance.CreditCardSecret) error
	UpdateCreditCardSecret(ctx context.Context, s *secretInstance.CreditCardSecret) error
	GetCreditCardSecret(ctx context.Context, secretID string) (*secretInstance.CreditCardSecret, error)
	ListCreditCardSecrets(ctx context.Context, accountID string) ([]*secretInstance.CreditCardSecret, error)
	DeleteCreditCardSecret(ctx context.Context, secretID string) error

	AddFileSecret(ctx context.Context, s *secretInstance.FileSecret) error
	UpdateFileSecret(ctx context.Context, s *secretInstance.FileSecret) error
	GetFileSecret(ctx context.Context, secretID string) (*secretInstance.FileSecret, error)
	ListFileSecrets(ctx context.Context, accountID string) ([]*secretInstance.FileSecret, error)
	DeleteFileSecret(ctx context.Context, secretID string) error

	MoveFileSecretToDeletionQueue(ctx context.Context, objectId string, ts time.Time) error
	ListFileDeletionQueue(ctx context.Context) ([]*secretInstance.DeletionQueueElement, error)
	DeleteFileSecretFromDeletionQueue(ctx context.Context, secretID string) error
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

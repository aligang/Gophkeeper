package handler

import (
	"context"
	"errors"
	account2 "github.com/aligang/Gophkeeper/pkg/common/account"
	instance2 "github.com/aligang/Gophkeeper/pkg/common/account/instance"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	"github.com/aligang/Gophkeeper/pkg/common/token/instance"
	"github.com/aligang/Gophkeeper/pkg/server/config"
	"github.com/aligang/Gophkeeper/pkg/server/encryption"
	"github.com/aligang/Gophkeeper/pkg/server/repository"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type GrpcHandler struct {
	account2.UnimplementedAccountServiceServer
	storage repository.Storage
	conf    *config.Config
}

func New(storage repository.Storage, config *config.Config) *GrpcHandler {
	return &GrpcHandler{
		storage: storage,
		conf:    config,
	}
}

func (h *GrpcHandler) Register(ctx context.Context, request *account2.RegisterRequest) (*account2.RegisterResponse, error) {

	logger := logging.Logger.GetSubLogger("AccountService", "Register")
	logger.Debug("Accepted new request")

	var err error
	id := uuid.New()
	instance := &instance2.Account{
		Id:                id.String(),
		Login:             request.Login,
		Password:          request.Password,
		EncryptionEnabled: true,
		CreatedAt:         time.Now(),
	}
	logger.Debug("generating encryption key...")
	instance.EncryptionKey, err = encryption.NewKey()
	if err != nil {

	}

	err = h.storage.WithinTransaction(ctx, func(tctx context.Context, tx *transaction.DBTransaction) error {
		_, terr := h.storage.GetAccountByLogin(tctx, request.Login, tx)
		if terr == nil {
			return account2.ErrRecordAlreadyExists
		}

		logger.Debug("generating encryption key succeeded")
		terr = h.storage.Register(
			tctx,
			instance,
			tx,
		)
		if terr != nil {
			return terr
		}
		return nil
	})

	switch {
	case errors.Is(err, account2.ErrRecordAlreadyExists):
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	case err != nil:
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	default:

	}
	logger.Debug("Sending response")
	return &account2.RegisterResponse{Account: instance2.ConvertAccountInstance(instance)}, nil
}

func (h *GrpcHandler) Authenticate(ctx context.Context, request *account2.AuthenticationRequest) (*account2.AuthenticationResponse, error) {

	logger := logging.Logger.GetSubLogger("AccountService", "Authenticate")
	logger.Debug("Received Authentication request")

	logger.Debug("Getting account information from sql")
	acc, err := h.storage.GetAccountByLogin(ctx, request.Login, nil)
	if err != nil {
		logger.Debug("account not found")
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	logger.Debug("Account information successfully retrieved")
	logger.Debug("Checking authentication information")
	if acc.Password != request.Password {
		logger.Debug("wrong  password")
		return nil, status.Errorf(codes.Unauthenticated, "Wrong password")
	}
	logger.Debug("Authentication information is valid")
	logger.Debug("Listing tokens for account: %s", acc.Id)
	accountTokens, err := h.storage.ListAccountTokens(ctx, acc.Id, nil)

	if err != nil {
		logger.Debug("Could not list tokens from database: %s", err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if len(accountTokens) == 0 {
		logger.Debug("No tokens were found within database")
	}
	var t *instance.Token
	if len(accountTokens) > 0 && accountTokens[0].IssuedAt.Add(
		time.Minute*time.Duration(h.conf.TokenRenewalTimeMinutes)).After(time.Now()) {
		logger.Debug("Using existing token")
		t = accountTokens[0]
	} else {
		logger.Debug("Creating new token record")
		t = instance.New(acc.Id)
		err = h.storage.WithinTransaction(ctx, func(tctx context.Context, tx *transaction.DBTransaction) error {
			logger.Debug("Adding new token record to sql id: %s", t.Id)
			return h.storage.AddToken(tctx, t, tx)
		})
		if err != nil {
			logger.Debug("Failed Issuing new token")
		}
		logger.Debug("New token record successfully created")
	}

	logger.Debug("Sending token to client")
	return &account2.AuthenticationResponse{Token: instance.ConvertTokenInstance(t)}, nil
}

package handler

import (
	"context"
	"errors"
	"github.com/aligang/Gophkeeper/internal/account"
	accountInstance "github.com/aligang/Gophkeeper/internal/account/instance"
	"github.com/aligang/Gophkeeper/internal/config"
	"github.com/aligang/Gophkeeper/internal/encryption"
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/aligang/Gophkeeper/internal/repository"
	"github.com/aligang/Gophkeeper/internal/repository/transaction"
	tokenInstance "github.com/aligang/Gophkeeper/internal/token/instance"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type GrpcHandler struct {
	account.UnimplementedAccountServiceServer
	storage repository.Storage
	conf    *config.ServerConfig
}

func New(storage repository.Storage, config *config.ServerConfig) *GrpcHandler {
	return &GrpcHandler{
		storage: storage,
		conf:    config,
	}
}

func (h *GrpcHandler) Register(ctx context.Context, request *account.RegisterRequest) (*account.RegisterResponse, error) {

	logger := logging.Logger.GetSubLogger("AccountService", "Register")
	logger.Debug("Accepted new request")

	var err error
	id := uuid.New()
	instance := &accountInstance.Account{
		Id:                id.String(),
		Login:             request.Login,
		Password:          request.Password,
		EncryptionEnabled: true,
	}
	logger.Debug("generating encryption key...")
	instance.EncryptionKey, err = encryption.NewKey()
	if err != nil {

	}

	err = h.storage.WithinTransaction(ctx, func(context.Context, *transaction.DBTransaction) error {
		_, terr := h.storage.GetAccountByLogin(ctx, request.Login)
		if terr == nil {
			return account.ErrRecordAlreadyExists
		}

		logger.Debug("generating encryption key succeeded")
		terr = h.storage.Register(
			ctx,
			instance,
		)
		if terr != nil {
			return terr
		}
		return nil
	})

	switch {
	case errors.Is(err, account.ErrRecordAlreadyExists):
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	case err != nil:
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	default:

	}
	logger.Debug("Sending response")
	return &account.RegisterResponse{Account: accountInstance.ConvertAccountInstance(instance)}, nil
}

func (h *GrpcHandler) Authenticate(ctx context.Context, request *account.AuthenticationRequest) (*account.AuthenticationResponse, error) {

	logger := logging.Logger.GetSubLogger("AccountService", "Authenticate")
	logger.Debug("Received Authentication request")

	logger.Debug("Getting account information from sql")
	acc, err := h.storage.GetAccountByLogin(ctx, request.Login)
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
	accountTokens, err := h.storage.ListAccountTokens(ctx, acc.Id)

	if err != nil {
		logger.Debug("Could not list tokens from database: %s", err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if len(accountTokens) == 0 {
		logger.Debug("No tokens were found within sql")
	}
	var t *tokenInstance.Token
	if len(accountTokens) > 0 && accountTokens[0].IssuedAt.Add(
		time.Minute*time.Duration(h.conf.TokenRenewalTimeMinutes)).After(time.Now()) {
		logger.Debug("Using existing token")
		t = accountTokens[0]
	} else {
		logger.Debug("Creating new token record")
		t = tokenInstance.New(acc.Id)
		err = h.storage.WithinTransaction(ctx, func(context.Context, *transaction.DBTransaction) error {
			logger.Debug("Adding new token record to sql id: %s", t.Id)
			return h.storage.AddToken(ctx, t)
		})
		if err != nil {
			logger.Debug("Failed Issuing new token")
		}
		logger.Debug("New token record successfully created")
	}

	logger.Debug("Sending token to client")
	return &account.AuthenticationResponse{Token: tokenInstance.ConvertTokenInstance(t)}, nil
}

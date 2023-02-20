package handler

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	secret "github.com/aligang/Gophkeeper/pkg/common/secret"
	"github.com/aligang/Gophkeeper/pkg/common/secret/instance"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

func (h *GrpcHandler) Create(ctx context.Context, req *secret.CreateSecretRequest) (*secret.SecretDescription, error) {
	logger := logging.Logger.GetSubLogger("handler", "Create")
	logger.Debug("Processing Create Secret Request")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Debug("Request missing metadata")
		return nil, status.Errorf(codes.FailedPrecondition, "Request missing metadata")
	}
	var accountID string
	values := md.Get("account_id")

	if len(values) == 0 {
		logger.Debug("Request missing account id information")
		return nil, status.Errorf(codes.Unauthenticated, "Request missing account id information")
	}
	accountID = values[0]
	logger.Debug("account id is: %s", accountID)

	desc := &secret.SecretDescription{}
	var err error
	secretId := uuid.New().String()

	logger.Info("Saving secret %s for account %s", secretId, accountID)
	switch req.Secret.(type) {
	case *secret.CreateSecretRequest_Text:
		s := &instance.TextSecret{
			BaseSecret: instance.BaseSecret{
				Id:        secretId,
				AccountId: accountID,
				CreatedAt: time.Now(),
			},
			Text: req.GetText().GetData(),
		}

		err = h.storage.WithinTransaction(ctx, func(ctx context.Context, tx *transaction.DBTransaction) error {
			return h.storage.AddTextSecret(ctx, s, tx)
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		desc.Id = s.Id
		desc.SecretType = secret.SecretType_TEXT
		desc.CreatedAt = s.CreatedAt.Format(time.RFC3339)
	case *secret.CreateSecretRequest_LoginPassword:
		s := &instance.LoginPasswordSecret{
			BaseSecret: instance.BaseSecret{
				Id:        secretId,
				AccountId: accountID,
				CreatedAt: time.Now(),
			},
			Login:    req.GetLoginPassword().GetLogin(),
			Password: req.GetLoginPassword().GetPassword(),
		}
		err = h.storage.WithinTransaction(ctx, func(ctx context.Context, tx *transaction.DBTransaction) error {
			return h.storage.AddLoginPasswordSecret(ctx, s, tx)
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		desc.Id = s.Id
		desc.SecretType = secret.SecretType_LOGIN_PASSWORD
		desc.CreatedAt = s.CreatedAt.Format(time.RFC3339)
	case *secret.CreateSecretRequest_CreditCard:
		s := &instance.CreditCardSecret{
			BaseSecret: instance.BaseSecret{
				Id:        secretId,
				AccountId: accountID,
				CreatedAt: time.Now(),
			},
			CardNumber: req.GetCreditCard().GetNumber(),
			CardHolder: req.GetCreditCard().GetCardholderName(),
			ValidTill:  req.GetCreditCard().GetValidTill(),
			Cvc:        req.GetCreditCard().GetCvc(),
		}
		err = h.storage.WithinTransaction(ctx, func(ctx context.Context, tx *transaction.DBTransaction) error {
			return h.storage.AddCreditCardSecret(ctx, s, tx)
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		desc.Id = s.Id
		desc.SecretType = secret.SecretType_CREDIT_CARD
		desc.CreatedAt = s.CreatedAt.Format(time.RFC3339)
	case *secret.CreateSecretRequest_File:
		s := &instance.FileSecret{
			BaseSecret: instance.BaseSecret{
				Id:        secretId,
				AccountId: accountID,
				CreatedAt: time.Now(),
			},
			ObjectId: uuid.New().String(),
		}
		logger.Debug("Saving file secret id=%s to object id=%s", s.Id, s.ObjectId)
		err = h.fileStorage.Save(ctx, s.ObjectId, req.GetFile().Data)
		if err != nil {
			logger.Debug("Failed to save file secret on file filerepository")
			return nil, status.Errorf(codes.Internal, "Failed to save file secret")
		}
		logger.Debug("File secret id=%s saved successfully ", s.Id)
		err = h.storage.WithinTransaction(ctx, func(ctx context.Context, tx *transaction.DBTransaction) error {
			return h.storage.AddFileSecret(ctx, s, tx)
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		desc.Id = s.Id
		desc.SecretType = secret.SecretType_FILE
		desc.CreatedAt = s.CreatedAt.Format(time.RFC3339)
	default:
		logger.Warn("Could not save secret %s for account %s", secretId, accountID)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	logger.Info("Secret %s for account %s successfully saved", secretId, accountID)
	logger.Debug("Finished Processing Create Secret Request, Sending Response")
	return desc, nil
}

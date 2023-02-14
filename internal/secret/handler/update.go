package handler

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/aligang/Gophkeeper/internal/repository/transaction"
	"github.com/aligang/Gophkeeper/internal/secret"
	"github.com/aligang/Gophkeeper/internal/secret/instance"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

func (h *GrpcHandler) Update(ctx context.Context, req *secret.UpdateSecretRequest) (*secret.SecretDescription, error) {
	logger := logging.Logger.GetSubLogger("handler", "Update")
	logger.Debug("Processing Update Secret Request")
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
	logger.Debug("%s", accountID)
	logger.Debug("Finished Processing Update Secret Request, Sending Response")

	var err error

	desc := &secret.SecretDescription{}

	switch req.Secret.(type) {
	case *secret.UpdateSecretRequest_Text:
		err = h.storage.WithinTransaction(ctx, func(context.Context, *transaction.DBTransaction) error {
			oldSecret, err := h.storage.GetTextSecret(ctx, req.Id)
			if err != nil {
				return err
			}
			err = CheckOwnership(oldSecret.AccountId, accountID)
			if err != nil {
				return err
			}
			s := &instance.TextSecret{
				BaseSecret: instance.BaseSecret{
					Id:         oldSecret.Id,
					AccountId:  oldSecret.AccountId,
					CreatedAt:  oldSecret.CreatedAt,
					ModifiedAt: time.Now(),
				},
				Text: req.GetText().GetData(),
			}
			desc.Id = s.Id
			desc.SecretType = secret.SecretType_TEXT
			desc.CreatedAt = s.CreatedAt.Format(time.RFC3339)
			desc.ModifiedAt = s.ModifiedAt.Format(time.RFC3339)

			return h.storage.UpdateTextSecret(ctx, s)
		})

	case *secret.UpdateSecretRequest_LoginPassword:
		err = h.storage.WithinTransaction(ctx, func(context.Context, *transaction.DBTransaction) error {
			oldSecret, err := h.storage.GetLoginPasswordSecret(ctx, req.Id)
			if err != nil {
				return err
			}
			err = CheckOwnership(oldSecret.AccountId, accountID)
			if err != nil {
				return err
			}
			s := &instance.LoginPasswordSecret{
				BaseSecret: instance.BaseSecret{
					Id:         oldSecret.Id,
					AccountId:  oldSecret.AccountId,
					CreatedAt:  oldSecret.CreatedAt,
					ModifiedAt: time.Now(),
				},
				Login:    req.GetLoginPassword().GetLogin(),
				Password: req.GetLoginPassword().GetPassword(),
			}
			desc.Id = s.Id
			desc.SecretType = secret.SecretType_LOGIN_PASSWORD
			desc.CreatedAt = s.CreatedAt.Format(time.RFC3339)
			desc.ModifiedAt = s.ModifiedAt.Format(time.RFC3339)
			return h.storage.UpdateLoginPasswordSecret(ctx, s)
		})
	case *secret.UpdateSecretRequest_CreditCard:
		err = h.storage.WithinTransaction(ctx, func(context.Context, *transaction.DBTransaction) error {
			oldSecret, err := h.storage.GetCreditCardSecret(ctx, req.Id)
			if err != nil {
				return err
			}
			err = CheckOwnership(oldSecret.AccountId, accountID)
			if err != nil {
				return err
			}
			s := &instance.CreditCardSecret{
				BaseSecret: instance.BaseSecret{
					Id:         oldSecret.Id,
					AccountId:  oldSecret.AccountId,
					CreatedAt:  oldSecret.CreatedAt,
					ModifiedAt: time.Now(),
				},
				Number:     req.GetCreditCard().GetNumber(),
				CardHolder: req.GetCreditCard().GetCardholderName(),
				ValidTill:  req.GetCreditCard().GetValidTill(),
				Cvc:        req.GetCreditCard().GetCvc(),
			}
			desc.Id = s.Id
			desc.SecretType = secret.SecretType_CREDIT_CARD
			desc.CreatedAt = s.CreatedAt.Format(time.RFC3339)
			desc.ModifiedAt = s.ModifiedAt.Format(time.RFC3339)
			return h.storage.UpdateCreditCardSecret(ctx, s)
		})
	case *secret.UpdateSecretRequest_File:
		logger.Debug("Running secret precheck...")
		oldSecret, err := h.storage.GetFileSecret(ctx, req.Id)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Secret not found")
		}
		logger.Debug("Succeed")
		logger.Debug("Checking permissions")
		err = CheckOwnership(oldSecret.AccountId, accountID)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "Access is prohibited")
		}
		logger.Debug("Succeed")
		newObjectId := uuid.New().String()
		err = h.fileStorage.Save(ctx, newObjectId, req.GetFile().Data)
		if err != nil {
			err = h.fileStorage.Delete(ctx, newObjectId)
			if err != nil {
				logger.Debug(
					"Error during erasing file=%s of secret=%s from filerepository",
					newObjectId, oldSecret.Id)
			}
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		s := &instance.FileSecret{
			BaseSecret: instance.BaseSecret{
				Id:         oldSecret.Id,
				AccountId:  oldSecret.AccountId,
				CreatedAt:  oldSecret.CreatedAt,
				ModifiedAt: time.Now(),
			},
			ObjectId: newObjectId,
		}
		logger.Debug("Updating object value to  %s", newObjectId)
		err = h.storage.WithinTransaction(ctx, func(context.Context, *transaction.DBTransaction) error {
			_, terr := h.storage.GetFileSecret(ctx, s.Id)
			if terr != nil {
				terr = h.fileStorage.Delete(ctx, s.ObjectId)
				if terr != nil {
					logger.Debug(
						"Error during erasing file=%s of secret=%s: Error during pre-check",
						s.ObjectId, s.Id)
				}
				return terr
			}

			terr = h.storage.UpdateFileSecret(ctx, s)
			if terr != nil {
				terr = h.fileStorage.Delete(ctx, s.ObjectId)
				if terr != nil {
					logger.Debug(
						"Error during erasing file=%s of secret=%s: Error during update record",
						s.ObjectId, s.Id)
				}
				return terr
			}

			terr = h.storage.MoveFileSecretToDeletionQueue(ctx, oldSecret.ObjectId, time.Now())
			if terr != nil {
				if terr != nil {
					logger.Debug(
						"Error during erasing file=%s of secret=%s: Error during update record",
						s.ObjectId, s.Id)
				}
				return terr
			}
			return nil
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		desc.Id = s.Id
		desc.SecretType = secret.SecretType_FILE
		desc.CreatedAt = s.CreatedAt.Format(time.RFC3339)
		desc.ModifiedAt = s.ModifiedAt.Format(time.RFC3339)
	default:
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Could not update secret")
	}

	return desc, nil
}

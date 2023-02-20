package handler

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	"github.com/aligang/Gophkeeper/pkg/common/secret"
	"github.com/aligang/Gophkeeper/pkg/common/secret/instance"
	"github.com/aligang/Gophkeeper/pkg/server/encryption"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
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
	logger.Info("Updating secret %s for account %s", req.Id, accountID)

	switch req.Secret.(type) {
	case *secret.UpdateSecretRequest_Text:
		err = h.storage.WithinTransaction(ctx, func(tCtx context.Context, tx *transaction.DBTransaction) error {
			oldSecret, err := h.storage.GetTextSecret(ctx, req.Id, tx)
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

			acc, terr := h.storage.GetAccountById(ctx, accountID, tx)
			if terr != nil {
				logger.Crit("Could not fetch account information from database for account %s", accountID)
				return terr
			}
			if h.isSecretEncryptionEnabled() {
				s, terr = encryption.EncryptTextSecret(s, acc.EncryptionKey)
				if err != nil {
					logger.Crit("Could not encrypt secret %s for account %s: %s", s.Id, accountID, err.Error())
					return terr
				}
			}

			desc.Id = s.Id
			desc.SecretType = secret.SecretType_TEXT
			desc.CreatedAt = s.CreatedAt.Format(time.RFC3339)
			desc.ModifiedAt = s.ModifiedAt.Format(time.RFC3339)

			return h.storage.UpdateTextSecret(ctx, s, tx)
		})

	case *secret.UpdateSecretRequest_LoginPassword:
		err = h.storage.WithinTransaction(ctx, func(tCtx context.Context, tx *transaction.DBTransaction) error {
			oldSecret, err := h.storage.GetLoginPasswordSecret(ctx, req.Id, tx)
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

			acc, terr := h.storage.GetAccountById(ctx, accountID, tx)
			if terr != nil {
				logger.Crit("Could not fetch account information from database for account %s", accountID)
				return terr
			}
			if h.isSecretEncryptionEnabled() {
				s, terr = encryption.EncryptLoginPasswordSecret(s, acc.EncryptionKey)
				if err != nil {
					logger.Crit("Could not encrypt secret %s for account %s: %s", s.Id, accountID, err.Error())
					return terr
				}
			}

			desc.Id = s.Id
			desc.SecretType = secret.SecretType_LOGIN_PASSWORD
			desc.CreatedAt = s.CreatedAt.Format(time.RFC3339)
			desc.ModifiedAt = s.ModifiedAt.Format(time.RFC3339)
			return h.storage.UpdateLoginPasswordSecret(ctx, s, tx)
		})
	case *secret.UpdateSecretRequest_CreditCard:
		err = h.storage.WithinTransaction(ctx, func(tCtx context.Context, tx *transaction.DBTransaction) error {
			oldSecret, err := h.storage.GetCreditCardSecret(ctx, req.Id, tx)
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
				CardNumber: req.GetCreditCard().GetNumber(),
				CardHolder: req.GetCreditCard().GetCardholderName(),
				ValidTill:  req.GetCreditCard().GetValidTill(),
				Cvc:        req.GetCreditCard().GetCvc(),
			}

			acc, terr := h.storage.GetAccountById(ctx, accountID, tx)
			if terr != nil {
				logger.Crit("Could not fetch account information from database for account %s", accountID)
				return terr
			}
			if h.isSecretEncryptionEnabled() {
				s, terr = encryption.EncryptCreditCardSecret(s, acc.EncryptionKey)
				if err != nil {
					logger.Crit("Could not encrypt secret %s for account %s: %s", s.Id, accountID, err.Error())
					return terr
				}
			}

			desc.Id = s.Id
			desc.SecretType = secret.SecretType_CREDIT_CARD
			desc.CreatedAt = s.CreatedAt.Format(time.RFC3339)
			desc.ModifiedAt = s.ModifiedAt.Format(time.RFC3339)
			return h.storage.UpdateCreditCardSecret(ctx, s, tx)
		})
	case *secret.UpdateSecretRequest_File:
		logger.Debug("Running secret precheck...")
		oldSecret, err := h.storage.GetFileSecret(ctx, req.Id, nil)
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
		err = h.storage.WithinTransaction(ctx, func(tCtx context.Context, tx *transaction.DBTransaction) error {
			_, terr := h.storage.GetFileSecret(ctx, s.Id, tx)
			if terr != nil {
				terr = h.fileStorage.Delete(ctx, s.ObjectId)
				if terr != nil {
					logger.Debug(
						"Error during erasing file=%s of secret=%s: Error during pre-check",
						s.ObjectId, s.Id)
				}
				return terr
			}

			terr = h.storage.UpdateFileSecret(ctx, s, tx)
			if terr != nil {
				terr = h.fileStorage.Delete(ctx, s.ObjectId)
				if terr != nil {
					logger.Debug(
						"Error during erasing file=%s of secret=%s: Error during update record",
						s.ObjectId, s.Id)
				}
				return terr
			}

			terr = h.storage.MoveFileSecretToDeletionQueue(ctx, oldSecret.ObjectId, time.Now(), tx)
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
		logger.Warn("Updating secret %s for account %s failed", req.Id, accountID)
		return nil, status.Errorf(codes.Unavailable, "Could not update secret")
	}
	logger.Warn("Updating secret %s for account %s succeeded", req.Id, accountID)
	return desc, nil
}

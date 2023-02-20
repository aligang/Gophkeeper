package handler

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	"github.com/aligang/Gophkeeper/pkg/common/secret"
	"github.com/aligang/Gophkeeper/pkg/server/encryption"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (h *GrpcHandler) Get(ctx context.Context, req *secret.GetSecretRequest) (*secret.Secret, error) {
	logger := logging.Logger.GetSubLogger("handler", "Get")
	logger.Debug("Processing Get Secret Request")
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

	var resp *secret.Secret
	logger.Info("Fetching secret %s for account %s", req.Id, accountID)

	acc, err := h.storage.GetAccountById(ctx, accountID, nil)
	if err != nil {
		logger.Crit("Could not fetch account information from database for account %s", accountID)
		return nil, status.Errorf(codes.Internal, "Account data not found")
	}

	switch req.SecretType {
	case secret.SecretType_TEXT:
		s, err := h.storage.GetTextSecret(ctx, req.Id, nil)

		if err != nil {
			return nil, status.Errorf(codes.NotFound, "Secret not found")
		}
		if err = CheckOwnership(s.AccountId, accountID); err != nil {
			return nil, status.Errorf(codes.Unavailable, "Access prohibited")
		}
		if h.isSecretEncryptionEnabled() {
			s, err = encryption.DecryptTextSecret(s, acc.EncryptionKey)
			if err != nil {
				logger.Crit("Could not decrypt secret %s for account %s: %s", req.Id, accountID, err.Error())
				return nil, status.Errorf(codes.Internal, "Could not decrypt secret")
			}
		}
		resp = convertTextSecretInstance(s)
	case secret.SecretType_LOGIN_PASSWORD:
		s, err := h.storage.GetLoginPasswordSecret(ctx, req.Id, nil)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "Secret not found")
		}
		if err = CheckOwnership(s.AccountId, accountID); err != nil {
			return nil, status.Errorf(codes.Unavailable, "Access prohibited")
		}
		if h.isSecretEncryptionEnabled() {
			s, err = encryption.DecryptLoginPasswordSecret(s, acc.EncryptionKey)
			if err != nil {
				logger.Crit("Could not decrypt secret %s for account %s: %s", req.Id, accountID, err.Error())
				return nil, status.Errorf(codes.Internal, "Could not decrypt secret")
			}
		}
		resp = convertLoginPasswordSecretInstance(s)
	case secret.SecretType_CREDIT_CARD:
		s, err := h.storage.GetCreditCardSecret(ctx, req.Id, nil)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "Secret not found")
		}
		if err = CheckOwnership(s.AccountId, accountID); err != nil {
			return nil, status.Errorf(codes.Unavailable, "Access prohibited")
		}
		if h.isSecretEncryptionEnabled() {
			s, err = encryption.DecryptCreditCardSecret(s, acc.EncryptionKey)
			if err != nil {
				logger.Crit("Could not decrypt secret %s for account %s: %s", req.Id, accountID, err.Error())
				return nil, status.Errorf(codes.Internal, "Could not decrypt secret")
			}
		}
		resp = convertCreditCardSecretInstance(s)
	case secret.SecretType_FILE:
		s, err := h.storage.GetFileSecret(ctx, req.Id, nil)

		if err != nil {
			logger.Debug("Secret record not found")
			return nil, status.Errorf(codes.NotFound, "Secret not found")
		}
		logger.Debug("Checking access permissions... ")
		if err = CheckOwnership(s.AccountId, accountID); err != nil {
			logger.Debug("failed ")
			return nil, status.Errorf(codes.Unavailable, "Access prohibited")
		}
		f, err := h.fileStorage.Read(ctx, s.ObjectId)
		if err != nil {
			logger.Debug("Secret data not found")
			return nil, status.Errorf(codes.NotFound, "Secret data not found")
		}
		resp = convertFileSecretInstance(s)
		resp.GetFile().Data = f
	default:
		logger.Debug("Unsupported secret type")
		return nil, status.Errorf(codes.InvalidArgument, "Unsupported secret type")
	}
	logger.Info("secret %s for account %s successfully fetched", req.Id, accountID)
	logger.Debug("Request is processed, sending response")
	return resp, nil
}

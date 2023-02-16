package handler

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/aligang/Gophkeeper/internal/secret"
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
	switch req.SecretType {
	case secret.SecretType_TEXT:
		s, err := h.storage.GetTextSecret(ctx, req.Id, nil)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "Secret not found")
		}
		if err = CheckOwnership(s.AccountId, accountID); err != nil {
			return nil, status.Errorf(codes.Unavailable, "Access prohibited")
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
		resp = convertLoginPasswordSecretInstance(s)
	case secret.SecretType_CREDIT_CARD:
		s, err := h.storage.GetCreditCardSecret(ctx, req.Id, nil)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "Secret not found")
		}
		if err = CheckOwnership(s.AccountId, accountID); err != nil {
			return nil, status.Errorf(codes.Unavailable, "Access prohibited")
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
		//return &secret.Secret{
		//	Id:         s.Id,
		//	CreatedAt:  s.CreatedAt.Format(time.RFC3339),
		//	ModifiedAt: s.ModifiedAt.Format(time.RFC3339),
		//	Secret:     &secret.Secret_File{File: &secret.File{Data: f}},
		//}, nil
	default:
		logger.Debug("Unsupported secret type")
		return nil, status.Errorf(codes.InvalidArgument, "Unsupported secret type")
	}
	logger.Debug("Request is processed, sending response")
	return resp, nil
}

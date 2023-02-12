package handler

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/aligang/Gophkeeper/internal/repository/transaction"
	"github.com/aligang/Gophkeeper/internal/secret"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func (h *GrpcHandler) Delete(ctx context.Context, req *secret.DeleteSecretRequest) (*empty.Empty, error) {
	logger := logging.Logger.GetSubLogger("handler", "Delete")
	logger.Debug("Processing Delete Secret Request")
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
	logger.Debug("Finished Processing Delete Secret Request, Sending Response")

	var err error

	switch req.SecretType {
	case secret.SecretType_TEXT:
		err = h.storage.WithinTransaction(ctx, func(context.Context, *transaction.DBTransaction) error {
			s, terr := h.storage.GetTextSecret(ctx, req.Id)
			if terr != nil {
				return terr
			}
			err = CheckOwnership(s.AccountId, accountID)
			if terr != nil {
				return terr
			}
			return h.storage.DeleteTextSecret(ctx, req.Id)
		})

	case secret.SecretType_LOGIN_PASSWORD:
		err = h.storage.WithinTransaction(ctx, func(context.Context, *transaction.DBTransaction) error {
			s, terr := h.storage.GetLoginPasswordSecret(ctx, req.Id)
			if terr != nil {
				return terr
			}
			err = CheckOwnership(s.AccountId, accountID)
			if terr != nil {
				return terr
			}
			return h.storage.DeleteLoginPasswordSecret(ctx, req.Id)
		})
	case secret.SecretType_CREDIT_CARD:
		err = h.storage.WithinTransaction(ctx, func(context.Context, *transaction.DBTransaction) error {
			s, terr := h.storage.GetCreditCardSecret(ctx, req.Id)
			if terr != nil {
				return terr
			}
			err = CheckOwnership(s.AccountId, accountID)
			if terr != nil {
				return terr
			}
			return h.storage.DeleteCreditCardSecret(ctx, req.Id)
		})
	case secret.SecretType_FILE:
		err = h.storage.WithinTransaction(ctx, func(context.Context, *transaction.DBTransaction) error {
			s, terr := h.storage.GetFileSecret(ctx, req.Id)
			if terr != nil {
				return terr
			}
			terr = h.storage.DeleteFileSecret(ctx, req.Id)
			if terr != nil {
				return terr
			}
			terr = h.storage.MoveFileSecretToDeletionQueue(ctx, s.ObjectId, time.Now())
			if terr != nil {
				return terr
			}
			return nil
		})
	default:
		return nil, status.Errorf(codes.Internal, secret.ErrUnsupportedSecretType.Error())
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

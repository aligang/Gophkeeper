package handler

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	secret2 "github.com/aligang/Gophkeeper/pkg/common/secret"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func (h *GrpcHandler) Delete(ctx context.Context, req *secret2.DeleteSecretRequest) (*empty.Empty, error) {
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
	case secret2.SecretType_TEXT:
		err = h.storage.WithinTransaction(ctx, func(tCtx context.Context, tx *transaction.DBTransaction) error {
			s, terr := h.storage.GetTextSecret(tCtx, req.Id, tx)
			if terr != nil {
				return terr
			}
			err = CheckOwnership(s.AccountId, accountID)
			if terr != nil {
				return terr
			}
			return h.storage.DeleteTextSecret(tCtx, req.Id, tx)
		})

	case secret2.SecretType_LOGIN_PASSWORD:
		err = h.storage.WithinTransaction(ctx, func(tCtx context.Context, tx *transaction.DBTransaction) error {
			s, terr := h.storage.GetLoginPasswordSecret(tCtx, req.Id, tx)
			if terr != nil {
				return terr
			}
			err = CheckOwnership(s.AccountId, accountID)
			if terr != nil {
				return terr
			}
			return h.storage.DeleteLoginPasswordSecret(tCtx, req.Id, tx)
		})
	case secret2.SecretType_CREDIT_CARD:
		err = h.storage.WithinTransaction(ctx, func(tCtx context.Context, tx *transaction.DBTransaction) error {
			s, terr := h.storage.GetCreditCardSecret(tCtx, req.Id, tx)
			if terr != nil {
				return terr
			}
			err = CheckOwnership(s.AccountId, accountID)
			if terr != nil {
				return terr
			}
			return h.storage.DeleteCreditCardSecret(tCtx, req.Id, tx)
		})
	case secret2.SecretType_FILE:
		err = h.storage.WithinTransaction(ctx, func(tCtx context.Context, tx *transaction.DBTransaction) error {
			s, terr := h.storage.GetFileSecret(tCtx, req.Id, tx)
			if terr != nil {
				return terr
			}
			terr = h.storage.DeleteFileSecret(tCtx, req.Id, tx)
			if terr != nil {
				return terr
			}
			terr = h.storage.MoveFileSecretToDeletionQueue(tCtx, s.ObjectId, time.Now(), tx)
			if terr != nil {
				return terr
			}
			return nil
		})
	default:
		return nil, status.Errorf(codes.Internal, secret2.ErrUnsupportedSecretType.Error())
	}
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not delete secret")
	}
	return &emptypb.Empty{}, nil
}

package handler

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	secret2 "github.com/aligang/Gophkeeper/pkg/common/secret"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

func (h *GrpcHandler) List(ctx context.Context, req *secret2.ListSecretRequest) (*secret2.ListSecretResponse, error) {
	logger := logging.Logger.GetSubLogger("handler", "List")
	logger.Debug("Processing List Secret Request")
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
	logger.Debug("Finished Processing Create Secret Request, Sending Response")

	response := &secret2.ListSecretResponse{}

	textSecrets, err := h.storage.ListTextSecrets(ctx, accountID, nil)
	if err != nil {
		logger.Debug("Error during fetching text secrets")
		return nil, status.Errorf(codes.Internal, "Error during fetching Text secrets")
	}
	loginPasswordSecrets, err := h.storage.ListLoginPasswordSecrets(ctx, accountID, nil)
	if err != nil {
		logger.Debug("Error during fetching text secrets")
		return nil, status.Errorf(codes.Internal, "Error during fetching Text secrets")
	}
	creditCardSecrets, err := h.storage.ListCreditCardSecrets(ctx, accountID, nil)
	if err != nil {
		logger.Debug("Error during fetching text secrets")
		return nil, status.Errorf(codes.Internal, "Error during fetching Text secrets")
	}
	fileSecrets, err := h.storage.ListFileSecrets(ctx, accountID, nil)
	if err != nil {
		logger.Debug("Error during fetching text secrets")
		return nil, status.Errorf(codes.Internal, "Error during fetching File secrets")
	}

	for _, s := range textSecrets {
		response.Secrets = append(
			response.Secrets,
			&secret2.SecretDescription{
				Id:         s.Id,
				CreatedAt:  s.CreatedAt.Format(time.RFC3339),
				ModifiedAt: convertTime(s.ModifiedAt),
				SecretType: secret2.SecretType_TEXT,
			},
		)
	}

	for _, s := range loginPasswordSecrets {
		response.Secrets = append(
			response.Secrets,
			&secret2.SecretDescription{
				Id:         s.Id,
				CreatedAt:  s.CreatedAt.Format(time.RFC3339),
				ModifiedAt: convertTime(s.ModifiedAt),
				SecretType: secret2.SecretType_LOGIN_PASSWORD,
			},
		)
	}

	for _, s := range creditCardSecrets {
		response.Secrets = append(
			response.Secrets,
			&secret2.SecretDescription{
				Id:         s.Id,
				CreatedAt:  s.CreatedAt.Format(time.RFC3339),
				ModifiedAt: convertTime(s.ModifiedAt),
				SecretType: secret2.SecretType_CREDIT_CARD,
			},
		)
	}

	for _, s := range fileSecrets {
		response.Secrets = append(
			response.Secrets,
			&secret2.SecretDescription{
				Id:         s.Id,
				CreatedAt:  s.CreatedAt.Format(time.RFC3339),
				ModifiedAt: convertTime(s.ModifiedAt),
				SecretType: secret2.SecretType_FILE,
			},
		)
	}

	return response, nil
}
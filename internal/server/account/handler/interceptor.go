package handler

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

func (h *GrpcHandler) AuthInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger := logging.Logger.GetSubLogger("Auth", "Interceptor")
	logger.Debug("Starting request preprocessing")
	if info.FullMethod != "/account.AccountService/Register" &&
		info.FullMethod != "/account.AccountService/Authenticate" {
		logger.Debug("Checking Metadata")
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			logger.Debug("Could not parse request metadata fields")
			return nil, status.Errorf(codes.Unauthenticated, "Could not parse request metadata fields")
		}
		var token string
		values := md.Get("token")

		if len(values) == 0 {
			logger.Debug("Could not find authorization token within metadata")
			return nil, status.Errorf(codes.Unauthenticated,
				"Could not fiend authorization token within metadata")
		}
		token = values[0]
		logger.Debug("Getting token record using token value from header")
		tokenRecord, err := h.storage.GetToken(ctx, token, nil)
		if err != nil {
			logger.Debug("Token is invalid or expired")
			return nil, status.Errorf(codes.Unauthenticated, "Token is invalid or expired")
		}
		logger.Debug("token record successfully retrieved from sql")
		if tokenRecord.IssuedAt.Add(time.Minute * time.Duration(h.conf.TokenValidityTimeMinutes)).Before(time.Now()) {
			logger.Debug("Token is expired")
			return nil, status.Errorf(codes.Unauthenticated, "Token is expired")
		}
		accountID := tokenRecord.Owner
		logger.Debug("Resolved account from token: %s", accountID)
		logger.Debug("Enriching Metadata")
		md.Append("account_id", accountID)
		ctx = metadata.NewIncomingContext(ctx, metadata.New(map[string]string{"account_id": accountID}))
	}
	logger.Debug("Finished request preprocessing")

	return handler(ctx, req)
}

package loginpassword

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/client/pipeline"
	"github.com/aligang/Gophkeeper/internal/client/token/tokengetter"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	secret2 "github.com/aligang/Gophkeeper/internal/common/secret"
	"google.golang.org/grpc/metadata"
)

func Get(client secret2.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Get LoginPassword")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	get := cli.Secret.LoginPassword.Get
	req := &secret2.GetSecretRequest{Id: get.Id, SecretType: secret2.SecretType_LOGIN_PASSWORD}

	logger.Debug("Encoding token into Metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"token": token.TokenValue}))

	logger.Debug("Sending request...")
	s, err := client.Get(ctx, req)
	if err != nil {
		logger.Fatal("Failed to Get Secret: %s", err.Error())
	}
	s.ToStdout()
	logger.Debug("Finished Pipeline execution")
}

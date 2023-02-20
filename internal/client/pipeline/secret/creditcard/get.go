package loginpassword

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/client/pipeline"
	"github.com/aligang/Gophkeeper/internal/client/token/tokengetter"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	"github.com/aligang/Gophkeeper/internal/common/secret"
	"google.golang.org/grpc/metadata"
)

func Get(client secret.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Get Credit catd")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	get := cli.Secret.CreditCard.Get
	req := &secret.GetSecretRequest{Id: get.Id, SecretType: secret.SecretType_CREDIT_CARD}

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

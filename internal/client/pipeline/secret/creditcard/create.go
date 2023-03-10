package loginpassword

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/client/pipeline"
	"github.com/aligang/Gophkeeper/internal/client/token/tokengetter"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	"github.com/aligang/Gophkeeper/internal/common/secret"
	"google.golang.org/grpc/metadata"
)

func Create(client secret.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Create Credit Card secret")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	create := cli.Secret.CreditCard.Create
	req := &secret.CreateSecretRequest{
		Secret: &secret.CreateSecretRequest_CreditCard{
			CreditCard: &secret.CreditCard{
				Number:         create.CardNumber,
				CardholderName: create.CardHolder,
				ValidTill:      create.ValidTill,
				Cvc:            create.Cvc,
			},
		},
	}
	logger.Debug("Encoding token into Metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"token": token.TokenValue}))

	logger.Debug("Sending request...")
	desc, err := client.Create(ctx, req)
	if err != nil {
		logger.Fatal("Failed to Create Secret: %s", err.Error())
		return
	}
	desc.ToStdout()
	logger.Debug("Finished Pipeline execution")
}

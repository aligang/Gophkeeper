package loginpassword

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/client/token/tokengetter"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	secret2 "github.com/aligang/Gophkeeper/pkg/common/secret"
	"google.golang.org/grpc/metadata"
)

func Create(client secret2.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Create Credit Card secret")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	create := cli.Secret.CreditCard.Create
	req := &secret2.CreateSecretRequest{
		Secret: &secret2.CreateSecretRequest_CreditCard{
			CreditCard: &secret2.CreditCard{
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
		logger.Debug("Failed to Create Secret: %s", err.Error())
		return
	}
	desc.ToStdout()
	logger.Debug("Finished Pipeline execution")
}

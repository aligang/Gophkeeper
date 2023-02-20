package loginpassword

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/client/pipeline"
	"github.com/aligang/Gophkeeper/internal/client/token/tokengetter"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	secret2 "github.com/aligang/Gophkeeper/internal/common/secret"
	"google.golang.org/grpc/metadata"
)

func Update(client secret2.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Update Credit Card")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	update := cli.Secret.CreditCard.Update
	req := &secret2.UpdateSecretRequest{
		Secret: &secret2.UpdateSecretRequest_CreditCard{
			CreditCard: &secret2.CreditCard{
				Number:         update.CardNumber,
				CardholderName: update.CardHolder,
				ValidTill:      update.ValidTill,
				Cvc:            update.Cvc,
			},
		},
		Id: update.Id,
	}
	logger.Debug("Encoding token into Metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"token": token.TokenValue}))

	logger.Debug("Sending request...")
	desc, err := client.Update(ctx, req)
	if err != nil {
		logger.Fatal("Failed to Update Secret: %s", err.Error())
	}
	desc.ToStdout()
	logger.Debug("Finished Pipeline execution")
}

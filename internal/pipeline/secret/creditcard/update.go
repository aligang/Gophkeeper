package loginpassword

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/aligang/Gophkeeper/internal/pipeline"
	"github.com/aligang/Gophkeeper/internal/secret"
	"github.com/aligang/Gophkeeper/internal/token/tokengetter"
	"google.golang.org/grpc/metadata"
)

func Update(client secret.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Update Credit Card")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	update := cli.Secret.CreditCard.Update
	req := &secret.UpdateSecretRequest{
		Secret: &secret.UpdateSecretRequest_CreditCard{
			CreditCard: &secret.CreditCard{
				Number:         update.CardNumber,
				CardholderName: update.CardHolder,
				ValidTill:      update.ValidTill,
				Cvc:            update.Cvc,
			},
		},
		Id: update.Id,
	}
	logger.Debug("Encoding token into Metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"token": token}))

	logger.Debug("Sending request...")
	desc, err := client.Update(ctx, req)
	if err != nil {
		logger.Debug("Failed to Update Secret: %s", err.Error())
		return
	}
	desc.ToStdout()
	logger.Debug("Finished Pipeline execution")
}

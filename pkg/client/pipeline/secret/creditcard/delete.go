package loginpassword

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/client/token/tokengetter"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	secret2 "github.com/aligang/Gophkeeper/pkg/common/secret"
	"google.golang.org/grpc/metadata"
)

func Delete(client secret2.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Delete Credit Card")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	del := cli.Secret.CreditCard.Delete
	req := &secret2.DeleteSecretRequest{Id: del.Id, SecretType: secret2.SecretType_CREDIT_CARD}

	logger.Debug("Encoding token into Metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"token": token.TokenValue}))

	logger.Debug("Sending request...")
	_, err := client.Delete(ctx, req)
	if err != nil {
		logger.Debug("Failed to delete Secret: %s", err.Error())
		return
	}
	logger.Debug("Finished Pipeline execution")
}

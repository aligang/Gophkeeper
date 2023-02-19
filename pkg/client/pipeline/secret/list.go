package secret

import (
	"context"
	"fmt"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/logging"
	"github.com/aligang/Gophkeeper/pkg/secret"
	"github.com/aligang/Gophkeeper/pkg/token/tokengetter"
	"google.golang.org/grpc/metadata"
)

func List(client secret.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "List")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	req := &secret.ListSecretRequest{}
	logger.Debug("Encoding token into Metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"token": token.TokenValue}))

	logger.Debug("Sending request...")
	secrets, err := client.List(ctx, req)
	if err != nil {
		logger.Debug("Failed to List Secrets: %s", err.Error())
		return
	}
	fmt.Print("[")
	for _, secret := range secrets.Secrets {
		fmt.Println()
		secret.ToStdout()
	}
	fmt.Println("]")
	logger.Debug("Finished Pipeline execution")
}

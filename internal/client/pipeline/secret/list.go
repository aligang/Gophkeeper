package secret

import (
	"context"
	"fmt"
	"github.com/aligang/Gophkeeper/internal/client/pipeline"
	"github.com/aligang/Gophkeeper/internal/client/token/tokengetter"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	secret2 "github.com/aligang/Gophkeeper/internal/common/secret"
	"google.golang.org/grpc/metadata"
)

func List(client secret2.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "List")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	req := &secret2.ListSecretRequest{}
	logger.Debug("Encoding token into Metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"token": token.TokenValue}))

	logger.Debug("Sending request...")
	secrets, err := client.List(ctx, req)
	if err != nil {
		logger.Fatal("Failed to List Secrets: %s", err.Error())
	}
	fmt.Print("[")
	for _, secret := range secrets.Secrets {
		fmt.Println()
		secret.ToStdout()
	}
	fmt.Println("]")
	logger.Debug("Finished Pipeline execution")
}

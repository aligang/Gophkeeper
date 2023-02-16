package loginpassword

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/aligang/Gophkeeper/internal/pipeline"
	"github.com/aligang/Gophkeeper/internal/secret"
	"github.com/aligang/Gophkeeper/internal/token/tokengetter"
	"google.golang.org/grpc/metadata"
)

func Create(client secret.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Create Login Password secret")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	create := cli.Secret.LoginPassword.Create
	req := &secret.CreateSecretRequest{
		Secret: &secret.CreateSecretRequest_LoginPassword{
			LoginPassword: &secret.LoginPassword{
				Login:    create.Login,
				Password: create.Password,
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

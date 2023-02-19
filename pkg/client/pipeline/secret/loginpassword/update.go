package loginpassword

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/logging"
	"github.com/aligang/Gophkeeper/pkg/secret"
	"github.com/aligang/Gophkeeper/pkg/token/tokengetter"
	"google.golang.org/grpc/metadata"
)

func Update(client secret.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Update Login Password")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	update := cli.Secret.LoginPassword.Update
	req := &secret.UpdateSecretRequest{
		Secret: &secret.UpdateSecretRequest_LoginPassword{
			LoginPassword: &secret.LoginPassword{
				Login:    update.Login,
				Password: update.Password,
			},
		},
		Id: update.Id,
	}
	logger.Debug(update.Id)
	logger.Debug("Encoding token into Metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"token": token.TokenValue}))

	logger.Debug("Sending request...")
	desc, err := client.Update(ctx, req)
	if err != nil {
		logger.Debug("Failed to Update Secret: %s", err.Error())
		return
	}
	desc.ToStdout()
	logger.Debug("Finished Pipeline execution")
}
package loginpassword

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/client/token/tokengetter"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	secret2 "github.com/aligang/Gophkeeper/pkg/common/secret"
	"google.golang.org/grpc/metadata"
)

func Update(client secret2.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Update Login Password")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	update := cli.Secret.LoginPassword.Update
	req := &secret2.UpdateSecretRequest{
		Secret: &secret2.UpdateSecretRequest_LoginPassword{
			LoginPassword: &secret2.LoginPassword{
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
		logger.Fatal("Failed to Update Secret: %s", err.Error())
	}
	desc.ToStdout()
	logger.Debug("Finished Pipeline execution")
}

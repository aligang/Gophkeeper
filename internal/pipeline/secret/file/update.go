package file

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/aligang/Gophkeeper/internal/pipeline"
	"github.com/aligang/Gophkeeper/internal/repository/fs"
	"github.com/aligang/Gophkeeper/internal/secret"
	"github.com/aligang/Gophkeeper/internal/token/tokengetter"
	"google.golang.org/grpc/metadata"
)

func Update(client secret.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Update File")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	update := cli.Secret.File.Update
	data, err := fs.ReadFile(context.Background(), update.FilePath)

	if err != nil {
		logger.Debug("Could not find file: %s", update.FilePath)
		return
	}

	req := &secret.UpdateSecretRequest{
		Id: update.Id,
		Secret: &secret.UpdateSecretRequest_File{
			File: &secret.File{
				Data: data,
			},
		},
	}
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

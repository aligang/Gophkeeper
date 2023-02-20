package file

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/client/token/tokengetter"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	secret2 "github.com/aligang/Gophkeeper/pkg/common/secret"
	"github.com/aligang/Gophkeeper/pkg/server/repository/fs"
	"google.golang.org/grpc/metadata"
)

func Update(client secret2.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Update File")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	update := cli.Secret.File.Update
	data, err := fs.ReadFile(context.Background(), update.FilePath)

	if err != nil {
		logger.Fatal("Could not find file: %s", update.FilePath)

	}

	req := &secret2.UpdateSecretRequest{
		Id: update.Id,
		Secret: &secret2.UpdateSecretRequest_File{
			File: &secret2.File{
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

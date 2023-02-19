package file

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/logging"
	"github.com/aligang/Gophkeeper/pkg/secret"
	"github.com/aligang/Gophkeeper/pkg/server/repository/fs"
	"github.com/aligang/Gophkeeper/pkg/token/tokengetter"
	"google.golang.org/grpc/metadata"
)

func Upload(client secret.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Upload File")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	upload := cli.Secret.File.Upload
	data, err := fs.ReadFile(context.Background(), upload.FilePath)

	if err != nil {
		logger.Debug("Could not find file: %s", upload.FilePath)
		return
	}

	req := &secret.CreateSecretRequest{
		Secret: &secret.CreateSecretRequest_File{
			File: &secret.File{
				Data: data,
			},
		},
	}
	logger.Debug("Encoding token into Metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"token": token.TokenValue}))

	logger.Debug("Sending request...")
	desc, err := client.Create(ctx, req)
	if err != nil {
		logger.Debug("Failed to Upload Secret: %s", err.Error())
		return
	}
	desc.ToStdout()
	logger.Debug("Finished Pipeline execution")
}

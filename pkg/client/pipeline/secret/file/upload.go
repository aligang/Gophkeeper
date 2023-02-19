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

func Upload(client secret2.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Upload File")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	upload := cli.Secret.File.Upload
	data, err := fs.ReadFile(context.Background(), upload.FilePath)

	if err != nil {
		logger.Debug("Could not find file: %s", upload.FilePath)
		return
	}

	req := &secret2.CreateSecretRequest{
		Secret: &secret2.CreateSecretRequest_File{
			File: &secret2.File{
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

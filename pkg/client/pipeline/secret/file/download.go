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

func Download(client secret.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Download File")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	download := cli.Secret.File.Download
	req := &secret.GetSecretRequest{Id: download.Id, SecretType: secret.SecretType_FILE}

	logger.Debug("Encoding token into Metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"token": token.TokenValue}))

	logger.Debug("Sending request...")
	resp, err := client.Get(ctx, req)
	if err != nil {
		logger.Debug("Failed to Download Secret: %s", err.Error())
		return
	}
	data := resp.Secret.(*secret.Secret_File).File.Data
	err = fs.SaveFile(context.Background(), download.FilePath, data)
	if err != nil {
		logger.Debug("Failed to save file: %s", err.Error())
		return
	}
	logger.Debug("Finished Pipeline execution")
}

package file

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/client/pipeline"
	"github.com/aligang/Gophkeeper/internal/client/token/tokengetter"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	secret2 "github.com/aligang/Gophkeeper/internal/common/secret"
	"github.com/aligang/Gophkeeper/internal/server/repository/fs"
	"google.golang.org/grpc/metadata"
)

func Download(client secret2.SecretServiceClient, getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {
	logger := logging.Logger.GetSubLogger("client pipeline", "Download File")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()
	download := cli.Secret.File.Download
	req := &secret2.GetSecretRequest{Id: download.Id, SecretType: secret2.SecretType_FILE}

	logger.Debug("Encoding token into Metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"token": token.TokenValue}))

	logger.Debug("Sending request...")
	resp, err := client.Get(ctx, req)
	if err != nil {
		logger.Fatal("Failed to Download Secret: %s", err.Error())
	}
	data := resp.Secret.(*secret2.Secret_File).File.Data
	err = fs.SaveFile(context.Background(), download.FilePath, data)
	if err != nil {
		logger.Debug("Failed to save file: %s", err.Error())
		return
	}
	logger.Debug("Finished Pipeline execution")
}

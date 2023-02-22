package token

import (
	"github.com/aligang/Gophkeeper/internal/client/pipeline"
	"github.com/aligang/Gophkeeper/internal/client/token/tokengetter"
	"github.com/aligang/Gophkeeper/internal/common/logging"
)

func Get(getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {

	logger := logging.Logger.GetSubLogger("token", "Get pipelien")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()

	token.ToStdout()
	logger.Debug("Finished Pipeline execution")
}

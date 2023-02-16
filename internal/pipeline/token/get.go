package token

import (
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/aligang/Gophkeeper/internal/pipeline"
	"github.com/aligang/Gophkeeper/internal/token/tokengetter"
)

func Get(getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {

	logger := logging.Logger.GetSubLogger("token", "Get pipelien")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()

	token.ToStdout()
	logger.Debug("Finished Pipeline execution")
}

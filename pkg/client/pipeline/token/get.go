package token

import (
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/logging"
	"github.com/aligang/Gophkeeper/pkg/token/tokengetter"
)

func Get(getter *tokengetter.TokenGetter, cli *pipeline.PipelineInitTree) {

	logger := logging.Logger.GetSubLogger("token", "Get pipelien")
	logger.Debug("Starting Pipeline execution")

	token := getter.GetToken()

	token.ToStdout()
	logger.Debug("Finished Pipeline execution")
}

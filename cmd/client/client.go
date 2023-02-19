package main

import (
	"github.com/aligang/Gophkeeper/pkg/client/config"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline/dispatcher"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logging.Configure(os.Stdout, zerolog.DebugLevel)
	logging.Debug("Starting GophKeeper client")
	clientCfg := config.GetConfig()
	pipelineCfg := pipeline.GetClientPipelineParamsFromCli()
	dispatcher.Start(clientCfg, pipelineCfg)
}

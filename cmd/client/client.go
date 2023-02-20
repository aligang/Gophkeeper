package main

import (
	"github.com/aligang/Gophkeeper/internal/client/config"
	"github.com/aligang/Gophkeeper/internal/client/pipeline"
	"github.com/aligang/Gophkeeper/internal/client/pipeline/dispatcher"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	"os"
)

func main() {
	pipelineCfg := &pipeline.PipelineInitTree{}
	clientCfg := config.GetConfig(pipelineCfg)
	logging.Init(os.Stdout)
	logging.SetLogLevel(clientCfg.LogLevel)
	logging.Info("Starting GophKeeper client")

	dispatcher.RunPipeline(clientCfg, pipelineCfg)
	logging.Info("Shutting Down GophKeeper client")
}

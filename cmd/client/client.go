package main

import (
	"github.com/aligang/Gophkeeper/internal/config"
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/aligang/Gophkeeper/internal/pipeline/dispatcher"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logging.Configure(os.Stdout, zerolog.DebugLevel)
	logging.Debug("Starting GophKeeper client")
	clientCfg := config.GetClientConfig()
	pipelineCfg := config.GetClientPipelineConfigFromCli()
	dispatcher.Start(clientCfg, pipelineCfg)
	//fmt.Println(clientCfg)
}

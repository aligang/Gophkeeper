package config

import (
	"flag"
	"fmt"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	"os"
)

func getCliConfig(pipelineCfg *pipeline.PipelineInitTree) *Config {
	cfg := &Config{}
	var logLevel string
	flag.StringVar(&cfg.ConfigFile, "c", "", "configuration file location. Default : /etc//etc/gophkeeper/client.yaml")
	flag.StringVar(&cfg.ServerAddress, "a", "", "host to listen on")
	flag.StringVar(&cfg.Login, "l", "", "File Storage Path")
	flag.StringVar(&cfg.Password, "p", "", "Config File Path")
	flag.StringVar(&logLevel, "log-level", "", "Logging level")
	flag.StringVar(&cfg.CaCertPath, "ca-cert", "", "CA certificate File Path")
	flag.BoolVar(&cfg.EnableTlsEncryption, "enable-channel-encryption", false, "Enable TLS Encryption for channel")

	if pipelineCfg != nil {
		pipeline.GetPipeline(pipelineCfg, func() {
			fmt.Fprintf(os.Stderr, "./gophkeeper-cli'.\n")
			fmt.Fprintf(os.Stderr, "    options:'.\n")
			fmt.Fprintf(os.Stderr, "      -a 'server address'.\n")
			fmt.Fprintf(os.Stderr, "      -l 'login'.\n")
			fmt.Fprintf(os.Stderr, "      -p 'password'.\n")
			fmt.Fprintf(os.Stderr, "      -log-level 'CRITICAL|DEBUG|WARNING'.\n")
		})
	} else {
		flag.Parse()
	}
	cfg.LogLevel = logging.GetLogLevelFromString(logLevel)
	return cfg
}

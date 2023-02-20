package config

import (
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	"os"
)

func GetConfig(pipelineCfg *pipeline.PipelineInitTree) *Config {
	var cfg *Config
	cfg = getConfigFromEnv().merge(getCliConfig(pipelineCfg))
	if cfg.ConfigFile != "" {
		cfg = cfg.merge(getConfigFromYaml(cfg.ConfigFile))
	} else if _, err := os.ReadFile(DEFAULT_CLIENT_CONFIG_FILE_LOCATION); err == nil {
		cfg = cfg.merge(getConfigFromYaml(DEFAULT_CLIENT_CONFIG_FILE_LOCATION))
	}
	cfg = cfg.merge(getClientDefaultConfig())
	return cfg
}

func (c *Config) merge(another *Config) *Config {
	if c.ServerAddress == "" && another.ServerAddress != "" {
		c.ServerAddress = another.ServerAddress
	}
	if c.Login == "" && another.Login != "" {
		c.Login = another.Login
	}
	if c.Password == "" && another.Password != "" {
		c.Password = another.Password
	}
	if c.LogLevel == logging.LogLevel_LOGLEVEL_UNSPECIFIED && another.LogLevel != logging.LogLevel_LOGLEVEL_UNSPECIFIED {
		c.LogLevel = another.LogLevel
	}
	return c
}

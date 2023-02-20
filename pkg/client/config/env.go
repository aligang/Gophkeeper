package config

import (
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	"os"
)

func getConfigFromEnv() *Config {
	s := Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		Login:         os.Getenv("LOGIN"),
		Password:      os.Getenv("PASSWORD"),
		LogLevel:      logging.GetLogLevelFromString(os.Getenv("LOGLEVEL")),
	}

	return &s
}

package config

import (
	"os"
)

func getConfigFromEnv() *Config {
	s := Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		Login:         os.Getenv("LOGIN"),
		Password:      os.Getenv("PASSWORD"),
	}
	return &s
}

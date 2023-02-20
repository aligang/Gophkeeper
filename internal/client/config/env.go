package config

import (
	"github.com/aligang/Gophkeeper/internal/common/logging"
	"os"
	"strconv"
)

func getConfigFromEnv() *Config {
	EnableTlsEncryption, err := strconv.ParseBool(os.Getenv("ENABLE_CHANNEL_ENCRYPTION"))
	if err != nil {
		EnableTlsEncryption = false
	}

	s := Config{
		ServerAddress:       os.Getenv("SERVER_ADDRESS"),
		Login:               os.Getenv("LOGIN"),
		Password:            os.Getenv("PASSWORD"),
		LogLevel:            logging.GetLogLevelFromString(os.Getenv("LOGLEVEL")),
		CaCertPath:          os.Getenv("CA_CERT_PASS"),
		EnableTlsEncryption: EnableTlsEncryption,
	}

	return &s
}

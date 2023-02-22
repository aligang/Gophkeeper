package config

import "github.com/aligang/Gophkeeper/internal/common/logging"

const DEFAULT_CLIENT_CONFIG_FILE_LOCATION = "/etc/gophkeeper/client.yaml"

func getClientDefaultConfig() *Config {
	return &Config{
		ServerAddress:       "127.0.0.1:8080",
		Login:               "user",
		Password:            "password",
		LogLevel:            logging.LogLevel_CRITICAL,
		CaCertPath:          "",
		EnableTlsEncryption: false,
	}
}

package config

const DEFAULT_CLIENT_CONFIG_FILE_LOCATION = "/etc/gophkeeper/client.yaml"

func getServerDefaultConfig() *ServerConfig {
	return &ServerConfig{
		Address:                  "127.0.0.1:8080",
		RepositoryType:           RepositoryType_IN_MEMORY,
		FileStorage:              "/tmp",
		TokenValidityTimeMinutes: 2,
		TokenRenewalTimeMinutes:  1,
		FileStaleTimeMinutes:     1,
		SecretEncryptionEnabled:  false,
	}
}

func getClientDefaultConfig() *ClientConfig {
	return &ClientConfig{
		ServerAddress: "127.0.0.1:8080",
		Login:         "user",
		Password:      "password",
	}
}

package config

import (
	"os"
	"strconv"
)

func getServerConfigFromEnv() *ServerConfig {
	repoType := os.Getenv("REPOSITORY_TYPE")
	TokenValidityTime, err := strconv.ParseInt(os.Getenv("TOKEN_VALIDITY_TIME"), 10, 32)
	if err != nil {
		TokenValidityTime = -1
	}
	TokenRenewalTime, err := strconv.ParseInt(os.Getenv("TOKEN_RENEWAL_TIME"), 10, 32)
	if err != nil {
		TokenRenewalTime = -1
	}
	FileStaleTime, err := strconv.ParseInt(os.Getenv("FILE_STALE_TIME"), 10, 32)
	if err != nil {
		FileStaleTime = -1
	}
	EnableSecretEncryption, err := strconv.ParseBool(os.Getenv("ENABLE_SECRET_ENCRYPTION"))
	if err != nil {
		EnableSecretEncryption = false
	}

	s := ServerConfig{
		Address:                  os.Getenv("ADDRESS"),
		RepositoryType:           getRepoValueFromName(&repoType),
		FileStorage:              os.Getenv("FILE_STORAGE"),
		ConfigFile:               os.Getenv("CONFIG_FILE"),
		TokenValidityTimeMinutes: TokenValidityTime,
		TokenRenewalTimeMinutes:  TokenRenewalTime,
		FileStaleTimeMinutes:     FileStaleTime,
		SecretEncryptionEnabled:  EnableSecretEncryption,
	}
	if s.RepositoryType == RepositoryType_SQL {
		//s.OptionalDatabaseUri = &ServerConfig_DatabaseUri{DatabaseUri: os.Getenv("DATABASE_URI")}
		s.DatabaseDsn = os.Getenv("DATABASE_DSN")
	}
	return &s
}

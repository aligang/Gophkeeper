package config

import (
	"github.com/aligang/Gophkeeper/internal/common/logging"
	"os"
	"strconv"
)

func getConfigFromEnv() *Config {
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

	s := Config{
		Address:                  os.Getenv("ADDRESS"),
		RepositoryType:           getRepoValueFromName(&repoType),
		FileStorage:              os.Getenv("FILE_STORAGE"),
		ConfigFile:               os.Getenv("CONFIG_FILE"),
		TokenValidityTimeMinutes: TokenValidityTime,
		TokenRenewalTimeMinutes:  TokenRenewalTime,
		FileStaleTimeMinutes:     FileStaleTime,
		SecretEncryptionEnabled:  EnableSecretEncryption,
		LogLevel:                 logging.GetLogLevelFromString(os.Getenv("LOGLEVEL")),
		TlsCertPath:              os.Getenv("TLS_CERT_PATH"),
		TlsKeyPath:               os.Getenv("TLS_KEY_PATH"),
	}
	if s.RepositoryType == RepositoryType_SQL {
		//s.OptionalDatabaseUri = &Config_DatabaseUri{DatabaseUri: os.Getenv("DATABASE_URI")}
		s.DatabaseDsn = os.Getenv("DATABASE_DSN")
	}
	return &s
}

package config

import (
	"github.com/aligang/Gophkeeper/pkg/common/logging"
)

func GetConfig() *Config {
	var cfg *Config
	cfg = getConfigFromEnv().merge(getConfigFromCli())
	if cfg.ConfigFile != "" {
		cfg = cfg.merge(getConfigFromYaml(cfg.ConfigFile))
	}
	cfg = cfg.merge(getServerDefaultConfig())
	return cfg
}

func (c *Config) merge(another *Config) *Config {
	if c.Address == "" && another.Address != "" {
		c.Address = another.Address
	}
	if c.RepositoryType == RepositoryType_UNSPECIFIED && another.RepositoryType != RepositoryType_UNSPECIFIED {
		c.RepositoryType = another.RepositoryType
	}
	if c.DatabaseDsn == "" && another.DatabaseDsn != "" {
		c.DatabaseDsn = another.DatabaseDsn
	}
	if c.FileStorage == "" && another.FileStorage != "" {
		c.FileStorage = another.FileStorage
	}

	if c.TokenValidityTimeMinutes <= 0 && another.TokenValidityTimeMinutes > 0 {
		c.TokenValidityTimeMinutes = another.TokenValidityTimeMinutes
	}

	if c.TokenRenewalTimeMinutes <= 0 && another.TokenRenewalTimeMinutes > 0 {
		c.TokenRenewalTimeMinutes = another.TokenRenewalTimeMinutes
	}

	if c.FileStaleTimeMinutes <= 0 && another.FileStaleTimeMinutes > 0 {
		c.FileStaleTimeMinutes = another.FileStaleTimeMinutes
	}

	if c.SecretEncryptionEnabled == false && another.SecretEncryptionEnabled == true {
		c.SecretEncryptionEnabled = true
	}
	if c.LogLevel == logging.LogLevel_LOGLEVEL_UNSPECIFIED && another.LogLevel != logging.LogLevel_LOGLEVEL_UNSPECIFIED {
		c.LogLevel = another.LogLevel
	}

	if c.ConfigFile == "" && another.ConfigFile != "" {
		c.ConfigFile = another.ConfigFile
	}
	return c
}

func getRepoValueFromName(n *string) RepositoryType {
	if repoType, ok := RepositoryType_value[*n]; ok {
		return RepositoryType(repoType)
	}
	return RepositoryType_UNSPECIFIED
}

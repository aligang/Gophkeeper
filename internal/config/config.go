package config

import "os"

func GetServerConfig() *ServerConfig {
	var cfg *ServerConfig
	cfg = getServerConfigFromEnv().merge(getServerConfigFromCli())
	if cfg.ConfigFile != "" {
		cfg = cfg.merge(getServerConfigFromYaml(cfg.ConfigFile))
	}
	cfg = cfg.merge(getServerDefaultConfig())
	return cfg
}

func (c *ServerConfig) merge(another *ServerConfig) *ServerConfig {
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

	if c.ConfigFile == "" && another.ConfigFile != "" {
		c.ConfigFile = another.ConfigFile
	}
	return c
}

func GetClientConfig() *ClientConfig {
	var cfg *ClientConfig
	cfg = getClientConfigFromEnv()
	if cfg.ConfigFile != "" {
		cfg = cfg.merge(getClientConfigFromYaml(cfg.ConfigFile))
	} else if _, err := os.ReadFile(DEFAULT_CLIENT_CONFIG_FILE_LOCATION); err == nil {
		cfg = cfg.merge(getClientConfigFromYaml(DEFAULT_CLIENT_CONFIG_FILE_LOCATION))
	}
	cfg = cfg.merge(getClientDefaultConfig())
	return cfg
}

func (c *ClientConfig) merge(another *ClientConfig) *ClientConfig {
	if c.ServerAddress == "" && another.ServerAddress != "" {
		c.ServerAddress = another.ServerAddress
	}
	if c.Login == "" && another.Login != "" {
		c.Login = another.Login
	}
	if c.Password == "" && another.Password != "" {
		c.Password = another.Password
	}
	return c
}

func getRepoValueFromName(n *string) RepositoryType {
	if repoType, ok := RepositoryType_value[*n]; ok {
		return RepositoryType(repoType)
	}
	return RepositoryType_UNSPECIFIED
}

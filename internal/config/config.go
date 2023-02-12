package config

import "github.com/aligang/Gophkeeper/internal/pipeline"

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

	if c.ConfigFile == "" && another.ConfigFile != "" {
		c.ConfigFile = another.ConfigFile
	}
	return c
}

func GetClientConfig() (*ClientConfig, *pipeline.PipelineInitTree) {
	var cfg *ClientConfig
	cfg = getClientConfigFromEnv()
	cliCfg, pipelineTree := getClientConfigFromCli()
	cfg = cfg.merge(cliCfg)
	//if cfg.ConfigFile != "" {
	//	cfg = cfg.merge(getServerConfigFromYaml(cfg.ConfigFile))
	//}
	cfg = cfg.merge(getClientDefaultConfig())
	return cfg, pipelineTree
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

package config

import "os"

func GetConfig() *Config {
	var cfg *Config
	cfg = getConfigFromEnv()
	if cfg.ConfigFile != "" {
		cfg = cfg.merge(getConfigFromYaml(cfg.ConfigFile))
	} else if _, err := os.ReadFile(DEFAULT_CLIENT_CONFIG_FILE_LOCATION); err == nil {
		cfg = cfg.merge(getConfigFromYaml(DEFAULT_CLIENT_CONFIG_FILE_LOCATION))
	}
	cfg = cfg.merge(getClientDefaultConfig())
	return cfg
}

func (c *Config) merge(another *Config) *Config {
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

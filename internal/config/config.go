package config

var cfg *Config

type Config struct {
	Database DatabaseConfig
	Secrets  SecretsConfig
	Auth     AuthConfig
}

func newConfig() *Config {
	return &Config{
		Database: *newDatabaseConfig(),
		Secrets:  *newSecretsConfig(),
		Auth:     *newAuthConfig(),
	}
}

func Get() *Config {
	if cfg == nil {
		cfg = newConfig()
	}

	return cfg
}

package config

type SecretsConfig struct {
	Application ApplicationSecretsConfig
}

type ApplicationSecretsConfig struct {
	ApplicationSecretKey string
}

func newSecretsConfig() *SecretsConfig {
	return &SecretsConfig{
		Application: ApplicationSecretsConfig{
			ApplicationSecretKey: getStrEnv("POKER_SECRET_KEY"),
		},
	}
}

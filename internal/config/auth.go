package config

type AuthConfig struct {
	TokenLifespanHours int
}

func newAuthConfig() *AuthConfig {
	return &AuthConfig{
		TokenLifespanHours: getIntEnv("TOKEN_LIFESPAN_HOURS"),
	}
}

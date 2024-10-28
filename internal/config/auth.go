package config

type AuthConfig struct {
	JwtToken *JwtTokenConfig
}

type JwtTokenConfig struct {
	LifespanHours int
	Issuer        string
}

func newAuthConfig() *AuthConfig {
	return &AuthConfig{
		JwtToken: &JwtTokenConfig{
			LifespanHours: getIntEnv("JWT_TOKEN_LIFESPAN_HOURS"),
			Issuer:        getStrEnv("JWT_TOKEN_ISSUER"),
		},
	}
}

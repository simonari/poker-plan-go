package config

type AuthConfig struct {
	JwtTokenConfig *JwtTokenConfig
}

type JwtTokenConfig struct {
	LifespanHours int
	Issuer        string
}

func newAuthConfig() *AuthConfig {
	return &AuthConfig{
		JwtTokenConfig: &JwtTokenConfig{
			LifespanHours: getIntEnv("JWT_TOKEN_LIFESPAN_HOURS"),
			Issuer:        getStrEnv("JWT_TOKEN_ISSUER"),
		},
	}
}

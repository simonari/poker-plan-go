package config

import "fmt"

type DatabaseCredentialsConfig struct {
	USER     string
	PASSWORD string
}

type DatabaseConnectionConfig struct {
	HOST     string
	PORT     int
	DATABASE string
}

type DatabaseDetailsConfig struct {
	SSLMode  string
	Timezone string
}

func getSSLModeEnv() string {
	map_ := map[bool]string{
		true:  "enable",
		false: "disable",
	}

	value := getBoolEnv("POSTGRES_SSL_MODE")

	return map_[value]
}

type DatabaseConfig struct {
	Credentials DatabaseCredentialsConfig
	Connection  DatabaseConnectionConfig
	Details     DatabaseDetailsConfig
}

func (c *DatabaseConfig) MakeDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		c.Connection.HOST,
		c.Connection.PORT,
		c.Credentials.USER,
		c.Credentials.PASSWORD,
		c.Connection.DATABASE,
		c.Details.SSLMode,
		c.Details.Timezone,
	)
}

func newDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Credentials: DatabaseCredentialsConfig{
			USER:     getStrEnv("POSTGRES_USER"),
			PASSWORD: getStrEnv("POSTGRES_PASSWORD"),
		},
		Connection: DatabaseConnectionConfig{
			HOST:     getStrEnv("POSTGRES_HOST"),
			PORT:     getIntEnv("POSTGRES_PORT"),
			DATABASE: getStrEnv("POSTGRES_DATABASE"),
		},
		Details: DatabaseDetailsConfig{
			SSLMode:  getSSLModeEnv(),
			Timezone: getStrEnv("POSTGRES_TIMEZONE"),
		},
	}
}

package config

import (
	"log"
	"os"
	"strconv"
)

type BaseConfig struct{}

func (c *BaseConfig) getBoolEnv(name string) bool {
	value_as_string, is_present := os.LookupEnv(name)

	if !is_present {
		log.Fatalf("[!] ENV [%s] was not found", name)
	}

	value, err := strconv.ParseBool(value_as_string)

	if err != nil {
		log.Fatalf("[!] Error converting [%s] to bool type", name)
	}

	return value
}

func (c *BaseConfig) getStrEnv(name string) string {
	value, is_present := os.LookupEnv(name)

	if !is_present {
		log.Fatalf("[!] ENV [%s] was not found", name)
	}

	return value
}

func (c *BaseConfig) getIntEnv(name string) int {
	value_as_string, is_present := os.LookupEnv(name)

	if !is_present {
		log.Fatalf("[!] ENV [%s] was not found", name)
	}

	value, err := strconv.Atoi(value_as_string)

	if err != nil {
		log.Fatalf("[!] Error converting [%s] to int type", name)
	}

	return value
}

type DatabaseCredentialsConfig struct {
	BaseConfig
	USER     string
	PASSWORD string
}

func (c *DatabaseCredentialsConfig) New() *DatabaseCredentialsConfig {
	return &DatabaseCredentialsConfig{
		USER:     c.getStrEnv("POSTGRES_USER"),
		PASSWORD: c.getStrEnv("POSTGRES_PASSWORD"),
	}
}

type DatabaseConnectionConfig struct {
	BaseConfig
	HOST     string
	PORT     int
	DATABASE string
}

func (c *DatabaseConnectionConfig) New() *DatabaseConnectionConfig {
	return &DatabaseConnectionConfig{
		HOST:     c.getStrEnv("POSTGRES_HOST"),
		PORT:     c.getIntEnv("POSTGRES_PORT"),
		DATABASE: c.getStrEnv("POSTGRES_DATABASE"),
	}
}

type DatabaseDetailsConfig struct {
	BaseConfig
	SSLMode  bool
	Timezone string
}

func (c *DatabaseDetailsConfig) New() *DatabaseDetailsConfig {
	return &DatabaseDetailsConfig{
		SSLMode:  c.getBoolEnv("POSTGRES_SSL_MODE"),
		Timezone: c.getStrEnv("POSTGRES_TIMEZONE"),
	}
}

type DatabaseConfig struct {
	Credentials DatabaseCredentialsConfig
	Connection  DatabaseConnectionConfig
	Details     DatabaseDetailsConfig
}

func (c *DatabaseConfig) MakeDSN() {

}

type Config struct {
	Database DatabaseConfig
}

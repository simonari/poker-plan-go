package config

import (
	"log"
	"os"
	"strconv"
)

func getBoolEnv(name string) bool {
	valueAsString, isPresent := os.LookupEnv(name)

	if !isPresent {
		log.Fatalf("[!] ENV [%s] was not found", name)
	}

	value, err := strconv.ParseBool(valueAsString)

	if err != nil {
		log.Fatalf("[!] Error converting [%s] to bool type", name)
	}

	return value
}

func getStrEnv(name string) string {
	value, isPresent := os.LookupEnv(name)

	if !isPresent {
		log.Fatalf("[!] ENV [%s] was not found", name)
	}

	return value
}

func getIntEnv(name string) int {
	valueAsString, isPresent := os.LookupEnv(name)

	if !isPresent {
		log.Fatalf("[!] ENV [%s] was not found", name)
	}

	value, err := strconv.Atoi(valueAsString)

	if err != nil {
		log.Fatalf("[!] Error converting [%s] to int type", name)
	}

	return value
}

var cfg *Config

type Config struct {
	Database DatabaseConfig
	Secrets  SecretsConfig
}

func newConfig() *Config {
	return &Config{
		Database: *newDatabaseConfig(),
		Secrets:  *newSecretsConfig(),
	}
}

func Get() *Config {
	if cfg == nil {
		cfg = newConfig()
	}

	return cfg
}

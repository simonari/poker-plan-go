package config

import (
	"log"
	"os"
	"strconv"
)

type BaseConfig struct{}

func getBoolEnv(name string) bool {
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

func getStrEnv(name string) string {
	value, is_present := os.LookupEnv(name)

	if !is_present {
		log.Fatalf("[!] ENV [%s] was not found", name)
	}

	return value
}

func getIntEnv(name string) int {
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

type Config struct {
	Database DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		Database: *newDatabaseConfig(),
	}
}

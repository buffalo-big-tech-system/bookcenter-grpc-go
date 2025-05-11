package config

import (
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		Host:     getEnvVal("DB_HOST"),
		Port:     getEnvVal("DB_PORT"),
		User:     getEnvVal("DB_USER"),
		Password: getEnvVal("DB_PASSWORD"),
		Name:     getEnvVal("DB_NAME"),
	}
}

func getEnvVal(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	panic("missing required environment variable: " + key)
}

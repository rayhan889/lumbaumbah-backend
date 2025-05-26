package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host                   string
	User                   string
	Password               string
	DBName                 string
	Port                   string
	SSLMode                string
	// JWTExpirationsInSecond int64
	// JWTSecret              string
}

func (c *Config) FormatDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}

var Envs = intitConfig()

func intitConfig() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	return Config{
		Host:    getEnv("DB_HOST", "https://localhost"),
		User:    getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:  getEnv("DB_NAME", "postgres"),
		Port:    getEnv("DB_PORT", "5431"),
		SSLMode: getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
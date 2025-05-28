package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host                   string
	User                   string
	Password               string
	DBName                 string
	DBPort                 string
	SSLMode                string
	Port                   string
	JWTExpirationsInSecond int64
	JWTSecret              string
}

func (c *Config) FormatDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.DBPort, c.DBName, c.SSLMode)
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
		DBPort:    getEnv("DB_PORT", "5431"),
		SSLMode: getEnv("DB_SSLMODE", "disable"),
		Port: getEnv("PORT", "8080"),
		JWTExpirationsInSecond: getEnvInt("JWT_EXPIRATIONS_IN_SECOND", 3600*24*7),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err :=strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
package config

import (
	"log"
	"os"
)

type Config struct {
	Port      string
	DBURL     string
	JWTSecret string
}

func Load() *Config {
	cfg := &Config{
		Port:      getEnv("PORT", "8080"),
		DBURL:     getEnv("DB_URL", "postgres://postgres:postgres@localhost:5432/auth?sslmode=disable"),
		JWTSecret: getEnv("JWT_SECRET", "super-secret-change-me"),
	}

	if cfg.JWTSecret == "super-secret-change-me" {
		log.Println("WARNING: JWT_SECRET is using default value. Change it in production.")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

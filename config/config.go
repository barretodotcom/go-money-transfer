package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func Init() (Config, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return Config{}, ErrDatabaseURLNotFound
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{DatabaseURL: databaseUrl, Port: port}, nil
}

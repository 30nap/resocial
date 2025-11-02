package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName   string
	AppEnv    string
	Port      string
	DBUrl     string
	JWTSecret string
}

var Cfg Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found in configs/, using system environment variables")
	}

	Cfg = Config{
		AppName:   getEnv("APP_NAME", "ReSocial"),
		AppEnv:    getEnv("APP_ENV", "development"),
		Port:      getEnv("PORT", "8000"),
		DBUrl:     getEnv("DATABASE_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
	}

	if Cfg.DBUrl == "" {
		log.Fatal("DATABASE_URL is required")
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

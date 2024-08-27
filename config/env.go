package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", ":8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "abc123"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
		DBName:     getEnv("DB_NAME", "ecomm"),
	}
}

var Envs = InitConfig()

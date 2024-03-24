package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
}

var Env = initConfig()

func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Port: getEnvVariable("PORT", "8080"),

		DBUser:     getEnvVariable("DB_USER", "root"),
		DBPassword: getEnvVariable("DB_PASSWORD", "password"),
		DBHost:     fmt.Sprintf("%s:%s", getEnvVariable("DB_HOST", "localhost"), getEnvVariable("DB_PORT", "3306")),
		DBName:     getEnvVariable("DB_NAME", "go-basic-ecom"),
	}
}

func getEnvVariable(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

package config

import (
	"log"
	"os"
	"strings"

	"github.com/core-wallet/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string `validate:"required"`

	DBHost     string `validate:"required"`
	DBPort     string `validate:"required"`
	DBUser     string `validate:"required"`
	DBPassword string `validate:"required"`
	DBName     string `validate:"required"`

	RedisHost     string
	RedisPort     string
	RedisPassword string

	CORSAllowed []string `validate:"required,dive,required"`
	AppSecret   string
}

var AppConfig *Config

func LoadConfig() {

	utils.InfoLog("INITIALIZE - APP CONFIG")

	AppConfig = &Config{
		AppPort:    getEnv("APP_PORT", "11000"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "postgres"),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),

		CORSAllowed: strings.Split(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173"), ","),
		AppSecret:   getEnv("APP_SECRET", ""),
	}

	validate := validator.New()
	if err := validate.Struct(AppConfig); err != nil {
		log.Fatalf("invalid config: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	log.Printf("WARNING: Environment variable %s is not set. Using fallback value: %s", key, fallback)
	return fallback
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		utils.ErrorLog("failed to load .env file", err, false)
	}
}
